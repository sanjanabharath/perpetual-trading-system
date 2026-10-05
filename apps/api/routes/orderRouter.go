package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderHandler struct {
	pool *pgxpool.Pool
}

func NewOrderHandler(pool *pgxpool.Pool) *OrderHandler {
	return &OrderHandler{
		pool: pool,
	}
}

// POST /orders/create
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order map[string]interface{}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := AddToQueue(order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}

// POST /orders/cancel
func (h *OrderHandler) CancelOrder(c *gin.Context) {
	var order map[string]interface{}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	err := AddToQueue(order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cancel order",
		})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// GET /orders/:id
func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID := c.Param("id")

	rows, err := h.pool.Query(
		context.Background(),
		`
		SELECT 
			o.id,
			o.user_id,
			o.symbol,
			o.quantity,
			o.price,
			t.id,
			t.order_id,
			t.price,
			t.quantity
		FROM orders o
		LEFT JOIN trades t ON t.order_id = o.id
		WHERE o.user_id = $1
		`,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch orders",
		})
		return
	}

	defer rows.Close()

	// We'll build the response here
	orders := make([]map[string]interface{}, 0)

	for rows.Next() {
		var (
			orderID      string
			userID       string
			symbol       string
			orderQty     float64
			orderPrice   float64

			tradeID      *string
			tradeOrderID *string
			tradePrice   *float64
			tradeQty     *float64
		)

		err := rows.Scan(
			&orderID,
			&userID,
			&symbol,
			&orderQty,
			&orderPrice,
			&tradeID,
			&tradeOrderID,
			&tradePrice,
			&tradeQty,
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read order",
			})
			return
		}

		order := map[string]interface{}{
			"id":       orderID,
			"userId":   userID,
			"symbol":   symbol,
			"quantity": orderQty,
			"price":    orderPrice,
		}

		if tradeID != nil {
			order["trade"] = map[string]interface{}{
				"id":       *tradeID,
				"orderId":  *tradeOrderID,
				"price":    *tradePrice,
				"quantity": *tradeQty,
			}
		}

		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch orders",
		})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found",
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}