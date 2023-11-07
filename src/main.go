package main

import (
	"GoDynamoApiTemplate/src/middleware"
	"GoDynamoApiTemplate/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()
	r.Use(middleware.CheckContentType())
	routes.SetupRouter(r)
	err := r.Run(":" + os.Getenv("GO_PORT"))
	if err != nil {
		return
	}
}
