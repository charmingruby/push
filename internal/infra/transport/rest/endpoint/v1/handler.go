package v1

import (
	docs "github.com/charmingruby/push/docs"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/observability/prometheus_observability"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	prometheus.MustRegister(prometheus_observability.HttpRequests)
	prometheus.MustRegister(prometheus_observability.RequestDuration)
}

func NewHTTPHandler(
	router *gin.Engine,
	notificationService notification_usecase.NotificationServiceUseCase,
) *HTTPHandler {
	return &HTTPHandler{
		router:              router,
		notificationService: notificationService,
	}
}

type HTTPHandler struct {
	router              *gin.Engine
	notificationService notification_usecase.NotificationServiceUseCase
}

func (h *HTTPHandler) Register() {
	basePath := "/api/v1"
	v1 := h.router.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath
	{
		v1.GET("/welcome", welcomeEndpoint)
		v1.POST("/communication-channels", h.createCommunicationChannelEndpoint)
		v1.POST("/notifications", h.scheduleNotificationEndpoint)
		v1.GET("/notifications/:id", h.getNotificationEndpoint)
		v1.PATCH("/notifications/:id/cancel", h.cancelNotificationEndpoint)
	}

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	h.router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
