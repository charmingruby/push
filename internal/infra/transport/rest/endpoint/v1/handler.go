package v1

import (
	"github.com/charmingruby/push/docs"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/observability/metric"
	"github.com/charmingruby/push/internal/infra/observability/metric/endpoint_metric"
	"github.com/charmingruby/push/internal/infra/transport/rest/endpoint"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPHandler(
	router *gin.Engine,
	notificationService notification_usecase.NotificationServiceUseCase,
	metric *metric.Metric,
) *HTTPHandler {
	return &HTTPHandler{
		router:              router,
		metrics:             metric,
		notificationService: notificationService,
	}
}

type HTTPHandler struct {
	router              *gin.Engine
	metrics             *metric.Metric
	notificationService notification_usecase.NotificationServiceUseCase
}

func (h *HTTPHandler) Register() {
	handlers := h.initHandlers()

	basePath := "/api/v1"
	v1 := h.router.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath

	for _, handler := range handlers {
		endpointWithMetrics := endpoint_metric.NewEndpointMetricAdapter(
			h.metrics.EndpointMetrics,
			handler,
			"v1",
		)

		v1.Handle(handler.Verb(), handler.Pattern(), endpointWithMetrics.AdaptHandler)
	}

	// Prometheus
	promHandler := promhttp.HandlerFor(h.metrics.Registry, promhttp.HandlerOpts{})
	h.router.GET("/metrics", gin.WrapH(promHandler))

	// Swagger
	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (h *HTTPHandler) initHandlers() []endpoint.EndpointHandler {
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

	return handlers
}
