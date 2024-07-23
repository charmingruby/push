package v1

import (
	"github.com/charmingruby/push/docs"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/observability/metrics"
	"github.com/charmingruby/push/internal/infra/transport/rest/endpoint"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	prometheus.MustRegister(metrics.HttpRequests)
	prometheus.MustRegister(metrics.RequestDuration)
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
	handlers := []endpoint.EndpointHandler{}

	welcomeEndpoint := NewWelcomeEndpoint()
	handlers = append(handlers, welcomeEndpoint)

	createCommunicationChannelEndpoint := NewCreateCommunicationChannelEndpoint(
		h.notificationService,
	)
	handlers = append(handlers, createCommunicationChannelEndpoint)

	scheduleNotificationEndpoint := NewScheduleNotificationEndpoint(
		h.notificationService,
	)
	handlers = append(handlers, scheduleNotificationEndpoint)

	getNotificationEndpoint := NewGetNotificationEndpoint(
		h.notificationService,
	)
	handlers = append(handlers, getNotificationEndpoint)

	cancelNotificationEndpoint := NewCancelNotificationEndpoint(
		h.notificationService,
	)
	handlers = append(handlers, cancelNotificationEndpoint)

	basePath := "/api/v1"
	v1 := h.router.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath
	for _, handler := range handlers {
		endpointWithMetrics := metrics.NewEndpointMetricsAdapter(handler)
		v1.Handle(handler.Verb(), handler.Pattern(), endpointWithMetrics.Handle)
	}

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	h.router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
