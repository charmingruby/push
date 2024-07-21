package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmingruby/push/config"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/database/mongo_repository"
	"github.com/charmingruby/push/internal/infra/transport/rest"
	v1 "github.com/charmingruby/push/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/push/pkg/dispatcher"
	"github.com/charmingruby/push/pkg/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cron "gopkg.in/robfig/cron.v2"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("CONFIGURATION: .env file not found")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("CONFIGURATION: %s", err.Error()))
		os.Exit(1)
	}

	db, err := mongodb.NewMongoConnection(cfg.MongoConfig.URL, cfg.MongoConfig.Database)
	if err != nil {
		slog.Error(fmt.Sprintf("MONGO CONNECTION: %s", err.Error()))
		os.Exit(1)
	}

	router := gin.Default()

	communicationChannelRepo := mongo_repository.NewCommunicationChannelMongoRepository(db)
	notificationRepo := mongo_repository.NewNotificationsMongoRepository(db)
	dispatcher := dispatcher.NewSimulationDispatcher()

	notificationSvc := notification_usecase.NewNotificationUseCaseRegistry(
		notificationRepo,
		communicationChannelRepo,
		dispatcher,
	)

	v1.NewHandler(router, notificationSvc).Register()

	server := rest.NewServer(router, cfg.ServerConfig.Port)

	runCronJobs(notificationSvc)

	go func() {
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error(fmt.Sprintf("REST SERVER: %s", err.Error()))
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	slog.Info("HTTP Server interruption received!")

	ctx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Server.Shutdown(ctx); err != nil {
		slog.Error(fmt.Sprintf("GRACEFUL SHUTDOWN REST SERVER: %s", err.Error()))
		os.Exit(1)
	}

	slog.Info("Gracefully shutdown!")
}

func runCronJobs(notificationSvc *notification_usecase.NotificationUseCaseRegistry) {
	c := cron.New()

	c.AddFunc("@every 00h01m00s", func() {
		slog.Info("[NOTIFICATION CRON JOB STATUS] Running...")

		notificationsWithFailure, err := notificationSvc.CheckAndSendNotificationUseCase()
		if err != nil {
			slog.Error("[NOTIFICATION CRON JOB ERROR] " + err.Error())
			return
		}

		if len(notificationsWithFailure) != 0 {
			slog.Info("[NOTIFICATION CRON JOB FAILED NOTIFICATIONS] " + fmt.Sprintf("%v", notificationsWithFailure))
			return
		}

		slog.Info("[NOTIFICATION CRON JOB STATUS] No errors")
	})

	c.Start()
}
