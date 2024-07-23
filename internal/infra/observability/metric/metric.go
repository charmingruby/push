package metric

import (
	"github.com/charmingruby/push/internal/infra/observability/metric/endpoint_metric"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Metric struct {
	EndpointMetrics *endpoint_metric.EndpointMetricRegistry
}

func NewMetrics(reg *prometheus.Registry) *Metric {
	m := &Metric{
		EndpointMetrics: endpoint_metric.NewEndpointMetricRegistry(),
	}

	reg.MustRegister(m.EndpointMetrics.EndpointLatency)
	reg.MustRegister(m.EndpointMetrics.RequestCounter)
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	return m
}
