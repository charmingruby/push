package v1

import (
	docs "github.com/charmingruby/push/docs"
	"github.com/charmingruby/push/internal/domain/example/example_usecase"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHandler(router *gin.Engine, exampleService example_usecase.ExampleServiceContract) *Handler {
	return &Handler{
		router:         router,
		exampleService: exampleService,
	}
}

type Handler struct {
	router         *gin.Engine
	exampleService example_usecase.ExampleServiceContract
}

func (h *Handler) Register() {
	basePath := "/api/v1"
	v1 := h.router.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath
	{
		v1.GET("/welcome", welcomeEndpoint)
		v1.POST("/examples", h.createExampleEndpoint)
		v1.GET("/examples/:id", h.getExampleEndpoint)
	}

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
