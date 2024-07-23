package v1

import (
	"log/slog"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/charmingruby/push/internal/infra/transport/rest"
	"github.com/gin-gonic/gin"
)

func NewGetNotificationEndpoint(service notification_usecase.NotificationServiceUseCase) *GetNotificationEndpoint {
	return &GetNotificationEndpoint{
		name:    "get notification",
		verb:    http.MethodGet,
		pattern: "/notifications/:id",
		service: service,
	}
}

type GetNotificationEndpoint struct {
	name    string
	verb    string
	pattern string
	service notification_usecase.NotificationServiceUseCase
}

type GetNotificationResponse struct {
	Message string                            `json:"message"`
	Data    *notification_entity.Notification `json:"data"`
}

// GetNotification godoc
//
//	@Summary		Gets a notification
//	@Description	Gets a notification
//	@Tags			Notifications
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Get Notification Payload"
//	@Success		200	{object}	GetNotificationResponse
//	@Failure		404	{object}	rest.Response
//	@Failure		500	{object}	rest.Response
//	@Router			/notifications/{id} [get]
func (h *GetNotificationEndpoint) Handle(c *gin.Context) (*rest.Response, *rest.Response) {
	notificationID := c.Param("id")

	dto := notification_dto.GetNotificationDTO{
		NotificationID: notificationID,
	}

	notification, err := h.service.GetNotificationUseCase(dto)
	if err != nil {
		resourceNotFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			res := rest.NewResourceNotFoundError(c, resourceNotFoundErr)
			return nil, &res
		}

		slog.Error(err.Error())
		res := rest.NewInternalServerError(c)
		return nil, &res
	}

	res := rest.NewOkResponse(
		c,
		"notification found",
		notification,
	)
	return &res, nil
}

func (h *GetNotificationEndpoint) Verb() string {
	return h.verb
}

func (h *GetNotificationEndpoint) Pattern() string {
	return h.pattern
}

func (h *GetNotificationEndpoint) Name() string {
	return h.name
}
