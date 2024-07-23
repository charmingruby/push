package metrics

import (
	"github.com/charmingruby/push/internal/infra/transport/rest/endpoint"
	"github.com/gin-gonic/gin"
)

func NewEndpointMetricsAdapter(endpointHandler endpoint.EndpointHandler) *EndpointMetricsAdapter {
	return &EndpointMetricsAdapter{
		endpointHandler: endpointHandler,
	}
}

type EndpointMetricsAdapter struct {
	endpointHandler endpoint.EndpointHandler
}

func (a *EndpointMetricsAdapter) Handle(c *gin.Context) {
	a.endpointHandler.Handle(c)
}
