package endpoint

import (
	"github.com/charmingruby/push/internal/infra/transport/rest"
	"github.com/gin-gonic/gin"
)

type EndpointHandler interface {
	Name() string
	Verb() string
	Pattern() string
	Handle(c *gin.Context) (res *rest.Response, err *rest.Response)
}
