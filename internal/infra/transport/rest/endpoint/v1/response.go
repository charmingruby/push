package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewResponse(c *gin.Context, code int, data any, message string) {
	res := Response{
		Message: message,
		Data:    data,
		Code:    code,
	}
	c.JSON(code, res)
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Code    int    `json:"status_code"`
}

func NewCreatedResponse(c *gin.Context, entity string) {
	msg := entity + " created successfully"
	NewResponse(c, http.StatusCreated, nil, msg)
}

func NewOkResponse(c *gin.Context, msg string, data any) {
	NewResponse(c, http.StatusOK, data, msg)
}

func NewPayloadError(c *gin.Context, err error) {
	NewResponse(c, http.StatusBadRequest, nil, "Payload error: "+err.Error())
}

func NewEntityError(c *gin.Context, err error) {
	NewResponse(c, http.StatusUnprocessableEntity, nil, err.Error())
}

func NewBadRequestError(c *gin.Context, err error) {
	NewResponse(c, http.StatusBadRequest, nil, err.Error())
}

func NewResourceNotFoundError(c *gin.Context, err error) {
	NewResponse(c, http.StatusNotFound, nil, err.Error())
}

func NewInternalServerError(c *gin.Context, err error) {
	NewResponse(c, http.StatusInternalServerError, nil, err.Error())
}
