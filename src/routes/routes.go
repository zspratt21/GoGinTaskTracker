package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(router *gin.Engine) *gin.Engine {
	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})
	colors := router.Group("/colors")
	ColorRoutes(colors)

	return router
}
