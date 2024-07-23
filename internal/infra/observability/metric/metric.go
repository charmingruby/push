package metric

import (
	"github.com/charmingruby/push/internal/infra/observability/metric/endpoint_metric"
	"github.com/charmingruby/push/internal/infra/observability/metric/job_metric/notification_job_metric"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type Metric struct {
	Registry        *prometheus.Registry
	EndpointMetrics *endpoint_metric.EndpointMetricRegistry
	JobMetrics      *JobMetric
}

type JobMetric struct {
	NotificationMetric *notification_job_metric.NotificationJobMetricRegistry
}

func NewMetrics() *Metric {
	m := &Metric{
		EndpointMetrics: endpoint_metric.NewEndpointMetricRegistry(),
		JobMetrics: &JobMetric{
			NotificationMetric: notification_job_metric.NewNotificationJobMetricRegistry(),
		},
	}

	m.Registry = prometheus.NewRegistry()

	m.Registry.MustRegister(m.EndpointMetrics.EndpointLatency)
	m.Registry.MustRegister(m.EndpointMetrics.RequestCounter)
	m.Registry.MustRegister(m.JobMetrics.NotificationMetric.NotificationProcessingFailures)

	m.Registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	return m
}
