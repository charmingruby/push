package notification_job_metric

import "time"

const (
	// Labels
	date string = "date"

	// Names
	notificationFailures = "notification_failures"
)

type NotificationJobMetrics struct {
	Date                time.Time
	FailedNotifications int
}
