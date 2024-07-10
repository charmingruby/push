package v1

import (
	"github.com/gin-gonic/gin"
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
	NewOkResponse(c, "OK!", nil)
}
