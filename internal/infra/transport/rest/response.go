package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewResponse(c *gin.Context, code int, data any, message string) Response {
	res := Response{
		Message:    message,
		StatusCode: code,
		Data:       data,
	}
	c.JSON(code, res)

	return res
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Data       any    `json:"data,omitempty"`
}

func NewCreatedResponse(c *gin.Context, entity string) Response {
	msg := entity + " created successfully"
	res := NewResponse(c, http.StatusCreated, nil, msg)
	return res
}

func NewOkResponse(c *gin.Context, msg string, data any) Response {
	res := NewResponse(c, http.StatusOK, data, msg)
	return res
}

func NewPayloadError(c *gin.Context, err error) Response {
	res := NewResponse(c, http.StatusBadRequest, nil, "Payload error: "+err.Error())
	return res
}

func NewEntityError(c *gin.Context, err error) Response {
	res := NewResponse(c, http.StatusUnprocessableEntity, nil, err.Error())
	return res
}

func NewBadRequestError(c *gin.Context, err error) Response {
	res := NewResponse(c, http.StatusBadRequest, nil, err.Error())
	return res
}

func NewResourceNotFoundError(c *gin.Context, err error) Response {
	res := NewResponse(c, http.StatusNotFound, nil, err.Error())
	return res
}

func NewInternalServerError(c *gin.Context) Response {
	res := NewResponse(c, http.StatusInternalServerError, nil, "internal server error")
	return res
}

func NewConflicError(c *gin.Context, err error) Response {
	res := NewResponse(c, http.StatusConflict, nil, err.Error())
	return res
}
