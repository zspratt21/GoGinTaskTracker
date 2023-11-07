package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			log.Println("CheckContentType middleware triggered")
			if contentType := c.GetHeader("Content-Type"); contentType != "application/json" {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Content-Type must be 'application/json'",
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
