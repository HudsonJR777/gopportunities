package router

import (
	docs "github.com/HudsonJR777/gopportunities/docs"
	"github.com/HudsonJR777/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	{
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)

		v1.GET("/opening/:id", handler.ShowOpeningHandler)

		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)

		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)

	}

	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
