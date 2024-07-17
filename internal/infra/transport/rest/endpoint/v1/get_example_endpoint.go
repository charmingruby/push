package v1

import (
	"log/slog"

	"github.com/charmingruby/push/internal/core"
	"github.com/charmingruby/push/internal/domain/example/example_entity"
	"github.com/gin-gonic/gin"
)

type GetExampleResponse struct {
	Message string                  `json:"message"`
	Data    *example_entity.Example `json:"data"`
	Code    int                     `json:"status_code"`
}

// GetExample godoc
//
//	@Summary		Gets an example
//	@Description	Gets an example
//	@Tags			Examples
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Get Example Payload"
//	@Success		200	{object}	GetExampleResponse
//	@Failure		404	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/examples/{id} [get]
func (h *Handler) getExampleEndpoint(c *gin.Context) {
	exampleID := c.Param("id")

	example, err := h.exampleService.GetExample(exampleID)
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
		"example found",
		example,
	)
}
