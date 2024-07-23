package v1

import (
	"github.com/charmingruby/push/docs"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/observability/metric"
	"github.com/charmingruby/push/internal/infra/observability/metric/endpoint_metric"
	"github.com/charmingruby/push/internal/infra/transport/rest/endpoint"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPHandler(
	router *gin.Engine,
	notificationService notification_usecase.NotificationServiceUseCase,
) *HTTPHandler {
	reg := prometheus.NewRegistry()
	metric := metric.NewMetrics(reg)
	metricHandler := &MetricHandler{
		metric:   metric,
		registry: reg,
	}

	return &HTTPHandler{
		router:              router,
		metric:              metricHandler,
		notificationService: notificationService,
	}
}

type HTTPHandler struct {
	router              *gin.Engine
	metric              *MetricHandler
	notificationService notification_usecase.NotificationServiceUseCase
}

type MetricHandler struct {
	metric   *metric.Metric
	registry *prometheus.Registry
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
		endpointWithMetrics := endpoint_metric.NewEndpointMetricAdapter(handler)
		v1.Handle(handler.Verb(), handler.Pattern(), endpointWithMetrics.Handle)
	}

	promHandler := promhttp.HandlerFor(h.metric.registry, promhttp.HandlerOpts{})

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	h.router.GET("/metrics", gin.WrapH(promHandler))
}
