package routes

import (
	"api/handlers"
	"api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.New()

	router.Use(middleware.RequestId())
	router.Use(middleware.Logger())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/transactions", handlers.GetTransactions)
	router.POST("/transactions", handlers.CreateTransaction)
	return router
}
