package metric

import "github.com/prometheus/client_golang/prometheus"

type Metric struct {
	HttpRequestMetric *prometheus.CounterVec
	RequestDuration   *prometheus.HistogramVec
}

func NewMetrics(reg *prometheus.Registry) *Metric {
	m := &Metric{
		HttpRequestMetric: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_request_total",
				Help: "Number of requests",
			},
			[]string{"path"}),
		RequestDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Name: "http_request_duration_seconds",
				Help: "Duration of http request",
			},
			[]string{"path"}),
	}

	reg.MustRegister(m.HttpRequestMetric)
	reg.MustRegister(m.RequestDuration)

	return m
}
