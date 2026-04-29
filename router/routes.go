package router

import (
	"github.com/HudsonJR777/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("opening", handler.UpadateOpeningHandler)

		v1.GET("/opening", handler.ListOpeningsHandler)

	}
}
