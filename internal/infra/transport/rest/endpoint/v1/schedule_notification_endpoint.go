package v1

import (
	"log/slog"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/transport/rest"
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
//	@Success		201		{object}	rest.Response
//	@Failure		400		{object}	rest.Response
//	@Failure		404		{object}	rest.Response
//	@Failure		422		{object}	rest.Response
//	@Failure		500		{object}	rest.Response
//	@Router			/notifications [post]
func (h *ScheduleNotificationEndpoint) Handle(c *gin.Context) (*rest.Response, *rest.Response) {
	var req ScheduleNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res := rest.NewPayloadError(c, err)
		return nil, &res
	}

	dto := notification_dto.ScheduleNotificationDTO{
		RawDate:                req.RawDate,
		Destination:            req.Destination,
		CommunicationChannelID: req.CommunicationChannelID,
	}

	if err := h.service.ScheduleNotificationUseCase(dto); err != nil {
		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			res := rest.NewEntityError(c, validationErr)
			return nil, &res
		}

		notFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			res := rest.NewResourceNotFoundError(c, notFoundErr)
			return nil, &res
		}

		slog.Error(err.Error())
		res := rest.NewInternalServerError(c)
		return nil, &res
	}

	res := rest.NewCreatedResponse(c, "notification")
	return &res, nil
}

func (h *ScheduleNotificationEndpoint) Verb() string {
	return h.verb
}

func (h *ScheduleNotificationEndpoint) Pattern() string {
	return h.pattern
}

func (h *ScheduleNotificationEndpoint) Name() string {
	return h.name
}
