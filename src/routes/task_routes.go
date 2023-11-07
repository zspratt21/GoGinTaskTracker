package routes

import (
	"GoDynamoApiTemplate/src/controllers"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(rg *gin.RouterGroup) {
	rg.DELETE("/:id", controllers.DeleteTask)
	rg.GET("/:id", controllers.GetTask)
	rg.GET("", controllers.GetTasks)
	rg.POST("", controllers.AddTask)
	rg.PUT("/:id", controllers.UpdateTask)
}
