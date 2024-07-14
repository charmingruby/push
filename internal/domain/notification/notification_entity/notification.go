package notification_entity

import (
	"fmt"
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity/notification_value_object"
	"github.com/oklog/ulid/v2"
)

const (
	MAX_RETRIES = 3
)

func NewNotification(destination, rawDate, communicationChannelID string) (*Notification, error) {
	refDate := "2006-01-02 15:00"

	convDate, err := time.Parse(refDate, rawDate)
	if err != nil {
		return nil, core.NewValidationErr(
			fmt.Sprintf("unable to parse `%s` into date format: `%s`", rawDate, refDate),
		)
	}

	defaultStatus, err := notification_value_object.NewNotificationStatus(
		notification_value_object.NOTIFICATION_PENDING_STATUS,
	)
	if err != nil {
		return nil, core.NewValidationErr(err.Error())
	}

	e := Notification{
		ID:                     ulid.Make().String(),
		Destination:            destination,
		Date:                   convDate,
		Status:                 defaultStatus,
		CommunicationChannelID: communicationChannelID,
		Retries:                0,
		CreatedAt:              time.Now(),
	}

	if err := core.ValidateStruct(e); err != nil {
		return nil, err
	}

	return &e, nil
}

type Notification struct {
	ID                     string    `json:"id" validate:"required"`
	Destination            string    `json:"destination" validate:"required"`
	Date                   time.Time `json:"date" validate:"required"`
	Status                 string    `json:"status" validate:"required"`
	CommunicationChannelID string    `json:"communication_channel_id" validate:"required"`
	Retries                int       `json:"retries" `
	CreatedAt              time.Time `json:"created_at" validate:"required"`
}

func (n *Notification) Retry() error {
	if n.Retries == MAX_RETRIES {
		return core.NewValidationErr("max retries reached")
	}

	n.Retries += 1
	return nil
}

func (n *Notification) StatusSent() {
	sts, _ := notification_value_object.NewNotificationStatus(
		notification_value_object.NOTIFICATION_SENT_STATUS,
	)

	n.Status = sts
}

func (n *Notification) StatusPending() {
	sts, _ := notification_value_object.NewNotificationStatus(
		notification_value_object.NOTIFICATION_PENDING_STATUS,
	)

	n.Status = sts
}

func (n *Notification) StatusCanceled() {
	sts, _ := notification_value_object.NewNotificationStatus(
		notification_value_object.NOTIFICATION_CANCELED_STATUS,
	)

	n.Status = sts
}

func (n *Notification) StatusFailure() {
	sts, _ := notification_value_object.NewNotificationStatus(
		notification_value_object.NOTIFICATION_FAILURE_STATUS,
	)

	n.Status = sts
}

func (n *Notification) StatusRetrying() {
	sts, _ := notification_value_object.NewNotificationStatus(
		notification_value_object.NOTIFICATION_RETRYING_STATUS,
	)

	n.Status = sts
}
