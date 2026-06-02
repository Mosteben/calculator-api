package main

import (
	"github.com/Mosteben/calculator-api/handlers"
	"github.com/Mosteben/calculator-api/middleware"

	"github.com/Mosteben/calculator-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.POST("/login", handlers.Login)

protected := r.Group("/")
protected.Use(middleware.AuthMiddleware())

protected.POST("/calculate", handlers.Calculate)
protected.GET("/history", handlers.GetHistory)
	r.Run(":8080")
}