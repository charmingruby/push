package v1

import (
	"net/http"

	"github.com/charmingruby/push/internal/infra/observability/metrics"
	"github.com/charmingruby/push/internal/infra/transport/rest"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func NewWelcomeEndpoint() *WelcomeEndpoint {
	return &WelcomeEndpoint{
		name:    "health check",
		verb:    http.MethodGet,
		pattern: "/welcome",
	}
}

type WelcomeEndpoint struct {
	name    string
	verb    string
	pattern string
}

// Welcome godoc
//
//	@Summary		Health Check
//	@Description	Health Check
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	rest.Response
//	@Router			/welcome [get]
func (h *WelcomeEndpoint) Handle(c *gin.Context) (*rest.Response, *rest.Response) {
	timer := prometheus.NewTimer(
		metrics.RequestDuration.WithLabelValues(c.Request.URL.Path),
	)
	defer timer.ObserveDuration()
	metrics.HttpRequests.WithLabelValues(c.Request.URL.Path).Inc()

	res := rest.NewOkResponse(c, "OK!", nil)

	return &res, nil
}

func (h *WelcomeEndpoint) Verb() string {
	return h.verb
}

func (h *WelcomeEndpoint) Pattern() string {
	return h.pattern
}

func (h *WelcomeEndpoint) Name() string {
	return h.name
}
