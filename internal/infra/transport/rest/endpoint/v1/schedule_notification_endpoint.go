package v1

import (
	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/gin-gonic/gin"
)

type ScheduleNotificationRequest struct {
	Destination            string `json:"destination" binding:"required"`
	RawDate                string `json:"raw_date" binding:"required"`
	CommunicationChannelID string `json:"communication_channel_id" binding:"required"`
}

// ScheduleNotification godoc
//
//	@Summary		Schedules a notification
//	@Description	Schedules a notification
//	@Tags			Notifications
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ScheduleNotificationRequest	true	"Schedule Notification Payload"
//	@Success		201		{object}	Response
//	@Failure		400		{object}	Response
//	@Failure		404		{object}	Response
//	@Failure		422		{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/notifications [post]
func (h *Handler) scheduleNotificationEndpoint(c *gin.Context) {
	var req ScheduleNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		NewPayloadError(c, err)
		return
	}

	dto := notification_dto.ScheduleNotificationDTO{
		RawDate:                req.RawDate,
		Destination:            req.Destination,
		CommunicationChannelID: req.CommunicationChannelID,
	}

	if err := h.notificationService.ScheduleNotificationUseCase(dto); err != nil {
		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			NewEntityError(c, validationErr)
			return
		}

		notFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			NewResourceNotFoundError(c, notFoundErr)
			return
		}

		NewInternalServerError(c, err)
		return
	}

	NewCreatedResponse(c, "notification")
}
