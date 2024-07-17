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
	"github.com/charmingruby/push/internal/domain/example/example_usecase"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/database/mongo_repository"
	"github.com/charmingruby/push/internal/infra/transport/rest"
	v1 "github.com/charmingruby/push/internal/infra/transport/rest/endpoint/v1"
	"github.com/charmingruby/push/pkg/dispatcher"
	"github.com/charmingruby/push/pkg/mongodb"
	"github.com/charmingruby/push/test/inmemory"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
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

	initDependencies(router, db)

	server := rest.NewServer(router, cfg.ServerConfig.Port)

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

func initDependencies(router *gin.Engine, db *mongo.Database) {
	exampleRepo := inmemory.NewInMemoryExampleRepository()
	communicationChannelRepo := mongo_repository.NewCommunicationChannelMongoRepository(db)
	notificationRepo := inmemory.NewInMemoryNotificationRepository()

	dispatcher := dispatcher.NewSimulationDispatcher()

	exampleSvc := example_usecase.NewExampleService(
		exampleRepo,
	)

	notificationSvc := notification_usecase.NewNotificationUseCaseRegistry(
		notificationRepo,
		communicationChannelRepo,
		dispatcher,
	)

	v1.NewHandler(router, exampleSvc, notificationSvc).Register()
}
