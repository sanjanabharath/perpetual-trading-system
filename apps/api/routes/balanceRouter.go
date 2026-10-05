package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type BalanceHandler struct {
	pool *pgxpool.Pool
}

func NewBalanceHandler(pool *pgxpool.Pool) *BalanceHandler {
	return &BalanceHandler{
		pool: pool,
	}
}

func (h *BalanceHandler) GetBalance(c *gin.Context) {
	userID := c.Param("userId")

		var id string
		var name string

		err := h.pool.QueryRow(
			context.Background(),
			"SELECT id, name FROM users WHERE id = $1",
			userID,
		).Scan(&id, &name)

		if err != nil {
			if err == pgx.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "User not found",
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch balance",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Balance fetched successfully",
			"user": gin.H{
				"id":   id,
				"name": name,
			},
		})
	}
