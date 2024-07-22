package v1

import (
	"github.com/charmingruby/push/internal/infra/observability/prometheus_observability"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Welcome godoc
//
//	@Summary		Health Check
//	@Description	Health Check
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	Response
//	@Router			/welcome [get]
func welcomeEndpoint(c *gin.Context) {
	timer := prometheus.NewTimer(
		prometheus_observability.RequestDuration.WithLabelValues(c.Request.URL.Path),
	)
	defer timer.ObserveDuration()
	prometheus_observability.HttpRequests.WithLabelValues(c.Request.URL.Path).Inc()

	NewOkResponse(c, "OK!", nil)
}
