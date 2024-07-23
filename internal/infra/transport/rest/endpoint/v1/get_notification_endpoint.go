package v1

import (
	"log/slog"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_entity"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/gin-gonic/gin"
)

func NewGetNotificationEndpoint(service notification_usecase.NotificationServiceUseCase) *GetNotificationEndpoint {
	return &GetNotificationEndpoint{
		name:    "get notification",
		verb:    http.MethodGet,
		pattern: "/notifications",
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
//	@Failure		404	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/notifications/{id} [get]
func (h *HTTPHandler) getNotificationEndpoint(c *gin.Context) {
	notificationID := c.Param("id")

	dto := notification_dto.GetNotificationDTO{
		NotificationID: notificationID,
	}

	notification, err := h.notificationService.GetNotificationUseCase(dto)
	if err != nil {
		resourceNotFoundErr, ok := err.(*core.ErrNotFound)
		if ok {
			NewResourceNotFoundError(c, resourceNotFoundErr)
			return
		}

		slog.Error(err.Error())
		NewInternalServerError(c)
		return
	}

	NewOkResponse(
		c,
		"notification found",
		notification,
	)
}
