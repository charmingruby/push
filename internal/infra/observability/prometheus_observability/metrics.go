package prometheus_observability

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "Number of requests",
		},
		[]string{"path"})

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of http request",
		},
		[]string{"path"})
)
