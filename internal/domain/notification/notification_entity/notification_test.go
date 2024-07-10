package notification_entity

import (
	"fmt"
	"testing"
	"time"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity/notification_value_object"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func Test_NewNotification(t *testing.T) {
	t.Run("it should be able to create a notification", func(t *testing.T) {
		fakeCommunicationChannelID := ulid.Make().String()

		n, err := NewNotification(
			"email@email.com",
			"2024-07-10 15:00",
			fakeCommunicationChannelID,
		)

		sts, _ := notification_value_object.NewNotificationStatus(
			notification_value_object.NOTIFICATION_PENDING_STATUS,
		)

		assert.NoError(t, err)
		assert.Equal(t, n.Destination, "email@email.com")
		assert.Equal(t, n.CommunicationChannelID, fakeCommunicationChannelID)
		assert.Equal(t, n.Status, sts)
	})

	t.Run("it should be not able to convert a notification date successfully", func(t *testing.T) {
		rawDate := "2024-07-10 15:00"

		n, err := NewNotification(
			"email@email.com",
			rawDate,
			ulid.Make().String(),
		)

		assert.NoError(t, err)
		assert.Equal(t, n.Date.Day(), 10)
		assert.Equal(t, n.Date.Month(), time.Month(7))
		assert.Equal(t, n.Date.Year(), 2024)
		assert.Equal(t, n.Date.Hour(), 15)
		assert.Equal(t, n.Date.Minute(), 0)
	})

	t.Run("it should be not able to create a notification with invalid date format", func(t *testing.T) {
		refDate := "2006-01-02 15:00"
		rawDate := "2024-07-10 15"

		n, err := NewNotification(
			"email@email.com",
			rawDate,
			ulid.Make().String(),
		)

		assert.Error(t, err)
		assert.Nil(t, n)
		assert.Equal(t, err.Error(), core.NewValidationErr(
			fmt.Sprintf("unable to parse `%s` into date format: `%s`", rawDate, refDate),
		).Error())
	})

	t.Run("it should be not able to create a notification with struct errors", func(t *testing.T) {
		n, err := NewNotification(
			"",
			"2024-07-10 15:00",
			ulid.Make().String(),
		)

		assert.Error(t, err)
		assert.Nil(t, n)
		assert.Equal(t, err.Error(), core.NewValidationErr(
			core.ErrRequired("destination"),
		).Error())
	})
}

func Test_NotifcationStatus(t *testing.T) {
	t.Run("it should be able to update a notification status to SENT", func(t *testing.T) {
		n, err := NewNotification(
			"email@email.com",
			"2024-07-10 15:00",
			ulid.Make().String(),
		)
		assert.NoError(t, err)
		assert.NotNil(t, n)

		sts, err := notification_value_object.NewNotificationStatus(
			notification_value_object.NOTIFICATION_SENT_STATUS,
		)
		assert.NoError(t, err)

		n.StatusSent()

		assert.Equal(t, n.Status, sts)
	})

	t.Run("it should be able to update a notification status to CANCELED", func(t *testing.T) {
		n, err := NewNotification(
			"email@email.com",
			"2024-07-10 15:00",
			ulid.Make().String(),
		)
		assert.NoError(t, err)
		assert.NotNil(t, n)

		sts, err := notification_value_object.NewNotificationStatus(
			notification_value_object.NOTIFICATION_CANCELED_STATUS,
		)
		assert.NoError(t, err)

		n.StatusCanceled()

		assert.Equal(t, n.Status, sts)
	})

	t.Run("it should be able to update a notification status to PENDING", func(t *testing.T) {
		n, err := NewNotification(
			"email@email.com",
			"2024-07-10 15:00",
			ulid.Make().String(),
		)
		assert.NoError(t, err)
		assert.NotNil(t, n)

		sts, err := notification_value_object.NewNotificationStatus(
			notification_value_object.NOTIFICATION_PENDING_STATUS,
		)
		assert.NoError(t, err)

		n.StatusPending()

		assert.Equal(t, n.Status, sts)
	})
}
