package notification_entity

import (
	"fmt"
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity/notification_value_object"
	"github.com/oklog/ulid/v2"
)

func NewNotification(destination, rawDate, communicationChannelID string) (*Notification, error) {
	refDate := "2006-01-02 15:04:05"

	convDate, err := time.Parse(refDate, rawDate)
	if err != nil {
		return nil, core.NewValidationErr(
			fmt.Sprintf("unable to parse `%s` into date format: `%s`", rawDate, convDate.String()),
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
	CreatedAt              time.Time `json:"created_at" validate:"required" `
}
