package notification_entity

import (
	"testing"

	"github.com/charmingruby/push/internal/core"
	"github.com/stretchr/testify/assert"
)

func Test_NewCommunicationChannel(t *testing.T) {
	t.Run("it should be able to create a communication channel", func(t *testing.T) {
		c, err := NewCommunicationChannel(
			"SMS",
			"Mobile notification",
		)

		assert.NoError(t, err)
		assert.Equal(t, c.Name, "SMS")
		assert.Equal(t, c.Description, "Mobile notification")
	})

	t.Run("it should be not able to create a notification with struct errors", func(t *testing.T) {
		c, err := NewCommunicationChannel(
			"",
			"Mobile notification",
		)
		assert.Error(t, err)
		assert.Nil(t, c)
		assert.Equal(t, err.Error(), core.NewValidationErr(
			core.ErrRequired("name"),
		).Error())
	})
}
