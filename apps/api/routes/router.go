package routes

import (
	"api/handlers"
	"api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	router := gin.New()

	// Middleware
	router.Use(middleware.RequestId())
	router.Use(middleware.Logger())

	// Routes
	router.POST("/groups", handlers.CreateGroup)
	router.POST("/expenses", handlers.AddExpense)

	router.GET("/groups/:group_id/balances", handlers.GetBalances)
	router.GET("/groups/:group_id/settlements", handlers.GetSettlements)

	router.GET("/transactions", handlers.GetTransactions)
	router.POST("/transactions", handlers.CreateTransaction)

	return router
}