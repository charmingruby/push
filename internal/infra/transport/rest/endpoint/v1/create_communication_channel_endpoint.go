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
//	@Success		201		{object}	rest.Response
//	@Failure		400		{object}	rest.Response
//	@Failure		409		{object}	rest.Response
//	@Failure		422		{object}	rest.Response
//	@Failure		500		{object}	rest.Response
//	@Router			/communication-channels [post]
func (h *CreateCommunicationChannelEndpoint) Handle(c *gin.Context) (*rest.Response, *rest.Response) {
	var req CreateCommunicationChannelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		res := rest.NewPayloadError(c, err)
		return nil, &res
	}

	dto := notification_dto.CreateCommunicationChannelDTO{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := h.service.CreateCommunicationChannelUseCase(dto); err != nil {
		validationErr, ok := err.(*core.ErrValidation)
		if ok {
			res := rest.NewEntityError(c, validationErr)
			return nil, &res
		}

		conflictErr, ok := err.(*core.ErrConflict)
		if ok {
			res := rest.NewConflicError(c, conflictErr)
			return nil, &res
		}

		slog.Error(err.Error())
		res := rest.NewInternalServerError(c)
		return nil, &res
	}

	res := rest.NewCreatedResponse(c, "communication channel")
	return &res, nil
}

func (h *CreateCommunicationChannelEndpoint) Verb() string {
	return h.verb
}

func (h *CreateCommunicationChannelEndpoint) Pattern() string {
	return h.pattern
}

func (h *CreateCommunicationChannelEndpoint) Name() string {
	return h.name
}
