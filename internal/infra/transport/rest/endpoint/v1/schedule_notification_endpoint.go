package v1

import (
	"log/slog"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/gin-gonic/gin"
)

func NewScheduleNotificationEndpoint(service notification_usecase.NotificationServiceUseCase) *ScheduleNotificationEndpoint {
	return &ScheduleNotificationEndpoint{
		name:    "schedule notification",
		verb:    http.MethodPost,
		pattern: "/notifications",
		service: service,
	}
}

type ScheduleNotificationEndpoint struct {
	name    string
	verb    string
	pattern string
	service notification_usecase.NotificationServiceUseCase
}

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
func (h *HTTPHandler) scheduleNotificationEndpoint(c *gin.Context) {
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

		slog.Error(err.Error())
		NewInternalServerError(c)
		return
	}

	NewCreatedResponse(c, "notification")
}
