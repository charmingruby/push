package v1

import "github.com/gin-gonic/gin"

type EndpointHandler interface {
	Name() string
	Verb() string
	Pattern() string
	Handle(c *gin.Context) (res *Response, err *Response)
}
