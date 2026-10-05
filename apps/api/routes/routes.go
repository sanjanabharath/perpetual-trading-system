package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router *gin.Engine, pool *pgxpool.Pool) {

	// Create handlers
	orderHandler := NewOrderHandler(pool)
	positionHandler := NewPositionHandler(pool)
	balanceHandler := NewBalanceHandler(pool)
	depthHandler := NewDepthHandler(pool)

	// Order routes
	router.POST("/orders/create", orderHandler.CreateOrder)
	router.POST("/orders/cancel", orderHandler.CancelOrder)
	router.GET("/orders/:id", orderHandler.GetOrders)

	// Position
	router.GET("/position/:userId", positionHandler.GetPosition)

	// Balance
	router.GET("/balance/:userId", balanceHandler.GetBalance)

	// Depth
	router.GET("/depth", depthHandler.GetDepth)
}