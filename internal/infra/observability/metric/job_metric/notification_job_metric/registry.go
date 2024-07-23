package notification_job_metric

import "github.com/prometheus/client_golang/prometheus"

type NotificationJobMetricRegistry struct {
	NotificationProcessingFailures *prometheus.HistogramVec
}

func NewNotificationJobMetricRegistry() *NotificationJobMetricRegistry {
	return &NotificationJobMetricRegistry{
		NotificationProcessingFailures: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "job",
			Name:      notificationFailures,
			Help:      "Number of jobs that failed in a process",
		}, []string{date}),
	}
}

func (m *NotificationJobMetricRegistry) Execute(metrics *NotificationJobMetrics) {
	labels := map[string]string{
		date: metrics.Date.String(),
	}

	m.NotificationProcessingFailures.WithLabelValues(
		labels[date],
	).Observe(float64(metrics.FailedNotifications))
}
