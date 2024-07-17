package v1

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/gin-gonic/gin"
)

type CancelNotificationResponse struct {
	Message string                            `json:"message"`
	Data    *notification_entity.Notification `json:"data"`
}

// CancelNotification godoc
//
//	@Summary		Cancel notification
//	@Description	Cancel notification
//	@Tags			Notifications
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Cancel Notification Payload"
//	@Success		200	{object}	CancelNotificationResponse
//	@Failure		404	{object}	Response
//	@Failure		422	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/notifications/{id}/cancel [patch]
func (h *Handler) cancelNotificationEndpoint(c *gin.Context) {
	notificationID := c.Param("id")

	dto := notification_dto.CancelNotificationDTO{
		NotificationID: notificationID,
	}

	err := h.notificationService.CancelNotiticationUseCase(dto)
	if err != nil {
		resourceNotFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			NewResourceNotFoundError(c, resourceNotFoundErr)
			return
		}

		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			NewEntityError(c, validationErr)
			return
		}

		NewInternalServerError(c, err)
		return
	}

	NewOkResponse(
		c,
		"notification canceled successfully",
		nil,
	)
}
