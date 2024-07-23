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

func NewCancelNotificationEndpoint(service notification_usecase.NotificationServiceUseCase) *CancelNotificationEndpoint {
	return &CancelNotificationEndpoint{
		name:    "cancel notification",
		verb:    http.MethodPatch,
		pattern: "/notifications/:id/cancel",
		service: service,
	}
}

type CancelNotificationEndpoint struct {
	name    string
	verb    string
	pattern string
	service notification_usecase.NotificationServiceUseCase
}

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
//	@Failure		404	{object}	rest.Response
//	@Failure		422	{object}	rest.Response
//	@Failure		500	{object}	rest.Response
//	@Router			/notifications/{id}/cancel [patch]
func (h *CancelNotificationEndpoint) Handle(c *gin.Context) (*rest.Response, *rest.Response) {
	notificationID := c.Param("id")

	dto := notification_dto.CancelNotificationDTO{
		NotificationID: notificationID,
	}

	err := h.service.CancelNotiticationUseCase(dto)
	if err != nil {
		resourceNotFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			res := rest.NewResourceNotFoundError(c, resourceNotFoundErr)
			return nil, &res
		}

		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			res := rest.NewEntityError(c, validationErr)
			return nil, &res

		}

		slog.Error(err.Error())
		res := rest.NewInternalServerError(c)
		return nil, &res
	}

	res := rest.NewOkResponse(
		c,
		"notification canceled successfully",
		nil,
	)
	return &res, nil
}

func (h *CancelNotificationEndpoint) Verb() string {
	return h.verb
}

func (h *CancelNotificationEndpoint) Pattern() string {
	return h.pattern
}

func (h *CancelNotificationEndpoint) Name() string {
	return h.name
}
