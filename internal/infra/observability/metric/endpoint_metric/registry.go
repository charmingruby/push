package endpoint_metric

import "github.com/prometheus/client_golang/prometheus"

type EndpointMetricRegistry struct {
	RequestCounter  *prometheus.CounterVec
	EndpointLatency *prometheus.HistogramVec
}

func NewEndpointMetricRegistry() *EndpointMetricRegistry {
	r := &EndpointMetricRegistry{
		RequestCounter: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: "api",
				Name:      endpointRequestCounter,
				Help:      "Number of requests of an endpoint",
			},
			[]string{
				endpointUrl,
				verb,
				pattern,
				failed,
				error,
				responseCode,
				isAvailabilityError,
				isReliabilityError,
			},
		),
		EndpointLatency: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: "api",
				Name:      endpointRequestLatency,
			},
			[]string{endpointUrl},
		),
	}

	return r
}
