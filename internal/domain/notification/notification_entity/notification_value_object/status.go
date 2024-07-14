package notification_value_object

import "fmt"

const (
	NOTIFICATION_SENT_STATUS     = "sent"
	NOTIFICATION_PENDING_STATUS  = "pending"
	NOTIFICATION_FAILURE_STATUS  = "failure"
	NOTIFICATION_CANCELED_STATUS = "canceled"
)

func NewNotificationStatus(status string) (string, error) {
	sts := map[string]string{
		NOTIFICATION_SENT_STATUS:     "SENT",
		NOTIFICATION_PENDING_STATUS:  "PENDING",
		NOTIFICATION_CANCELED_STATUS: "CANCELED",
		NOTIFICATION_FAILURE_STATUS:  "FAILURE",
	}

	s, ok := sts[status]
	if !ok {
		return "", fmt.Errorf("invalid status")
	}

	return s, nil
}
