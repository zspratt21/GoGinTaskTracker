package routes

import (
	"GoDynamoApiTemplate/src/controllers"
	"github.com/gin-gonic/gin"
)

func ColorRoutes(rg *gin.RouterGroup) {
	rg.DELETE("/:name", controllers.DeleteColor)
	rg.GET("/:name", controllers.GetColor)
	rg.GET("", controllers.GetColors)
	rg.POST("", controllers.AddColor)
	rg.PUT("/:name", controllers.UpdateColor)
}
