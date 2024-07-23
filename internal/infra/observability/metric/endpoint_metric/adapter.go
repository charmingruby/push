package endpoint_metric

import (
	"fmt"
	"time"

	"github.com/charmingruby/push/internal/infra/transport/rest"
	"github.com/charmingruby/push/internal/infra/transport/rest/endpoint"
	"github.com/gin-gonic/gin"
)

const (
	baseUrl = "/api"
)

func NewEndpointMetricAdapter(
	metricsRegistry *EndpointMetricRegistry,
	endpointHandler endpoint.EndpointHandler,
	version string,
) *EndpointMetricsAdapter {
	return &EndpointMetricsAdapter{
		metricsRegistry: metricsRegistry,
		endpointHandler: endpointHandler,
		version:         version,
	}
}

type EndpointMetricsAdapter struct {
	version         string
	endpointHandler endpoint.EndpointHandler
	metricsRegistry *EndpointMetricRegistry
}

func (a *EndpointMetricsAdapter) AdaptHandler(c *gin.Context) {
	a.execute(c)
}

func (a *EndpointMetricsAdapter) execute(c *gin.Context) {
	start := time.Now()
	res, err := a.endpointHandler.Handle(c)
	latency := time.Since(start).Seconds()

	a.metrify(res, err, latency)
}

func (a *EndpointMetricsAdapter) metrify(
	res *rest.Response,
	err *rest.Response,
	latencyInSeconds float64,
) {
	pattern := fmt.Sprintf("%s/%s/%s",
		baseUrl,
		a.version,
		a.endpointHandler.Pattern(),
	)

	metrics := EndpointMetrics{
		Endpoint: a.endpointHandler.Name(),
		Verb:     a.endpointHandler.Verb(),
		Pattern:  pattern,
		Latency:  latencyInSeconds,
	}

	if err != nil {
		metrics.Failed = true
		metrics.Error = err.Message
		metrics.ResponseCode = err.StatusCode

		if err.StatusCode >= 500 {
			metrics.HasReliabilityError = false
			metrics.HasAvailabilityError = true
		} else {
			metrics.HasReliabilityError = true
			metrics.HasAvailabilityError = false
		}
	} else {
		metrics.Failed = false
		metrics.ResponseCode = res.StatusCode
	}

	a.sendEndpointMetrics(metrics)
}

func (a *EndpointMetricsAdapter) sendEndpointMetrics(metrics EndpointMetrics) {
	labels := map[string]string{
		endpointUrl:         metrics.Endpoint,
		verb:                metrics.Verb,
		pattern:             metrics.Pattern,
		failed:              fmt.Sprintf("%v", metrics.Failed),
		error:               metrics.Error,
		responseCode:        fmt.Sprintf("%d", metrics.ResponseCode),
		isAvailabilityError: fmt.Sprintf("%v", metrics.HasAvailabilityError),
		isReliabilityError:  fmt.Sprintf("%v", metrics.HasReliabilityError),
	}

	//a.metricsRegistry.EndpointLatency
	a.metricsRegistry.RequestCounter.WithLabelValues(
		labels[endpointUrl],
		labels[verb],
		labels[pattern],
		labels[failed],
		labels[error],
		labels[responseCode],
		labels[isAvailabilityError],
		labels[isReliabilityError],
	).Inc()
}
