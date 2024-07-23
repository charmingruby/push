package endpoint_metric

import (
	"github.com/charmingruby/push/internal/infra/transport/rest/endpoint"
	"github.com/gin-gonic/gin"
)

func NewEndpointMetricAdapter(endpointHandler endpoint.EndpointHandler) *EndpointMetricsAdapter {
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
