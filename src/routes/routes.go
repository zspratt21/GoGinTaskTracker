package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "welcome.html", nil)
	})

	tasks := router.Group("/api/Tasks")
	TaskRoutes(tasks)

	return router
}
