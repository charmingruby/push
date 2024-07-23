package v1

import (
	"log/slog"
	"net/http"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/notification/notification_dto"
	"github.com/charmingruby/push/internal/domain/notification/notification_usecase"
	"github.com/gin-gonic/gin"
)

func NewCreateCommunicationChannelEndpoint(
	service notification_usecase.NotificationServiceUseCase,
) *CreateCommunicationChannelEndpoint {
	return &CreateCommunicationChannelEndpoint{
		name:    "create communication channel",
		verb:    http.MethodPost,
		pattern: "/communication-channels",
		service: service,
	}
}

type CreateCommunicationChannelEndpoint struct {
	name    string
	verb    string
	pattern string
	service notification_usecase.NotificationServiceUseCase
}

type CreateCommunicationChannelRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

// CreateCommunicationChannel godoc
//
//	@Summary		Creates a communication channel
//	@Description	Creates a communication channel
//	@Tags			Communication Channel
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateCommunicationChannelRequest	true	"Create Communication Channel Payload"
//	@Success		201		{object}	Response
//	@Failure		400		{object}	Response
//	@Failure		409		{object}	Response
//	@Failure		422		{object}	Response
//	@Failure		500		{object}	Response
//	@Router			/communication-channels [post]
func (h *HTTPHandler) createCommunicationChannelEndpoint(c *gin.Context) {
	var req CreateCommunicationChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		NewPayloadError(c, err)
		return
	}

	dto := notification_dto.CreateCommunicationChannelDTO{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.notificationService.CreateCommunicationChannelUseCase(dto); err != nil {
		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			NewEntityError(c, validationErr)
			return
		}

		conflictErr, ok := err.(*core.ErrConflict)
		if ok {
			NewConflicError(c, conflictErr)
			return
		}

		slog.Error(err.Error())
		NewInternalServerError(c)
		return
	}

	NewCreatedResponse(c, "communication channel")
}
