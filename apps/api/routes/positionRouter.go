package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PositionHandler struct {
	pool *pgxpool.Pool
}

func NewPositionHandler(pool *pgxpool.Pool) *PositionHandler {
	return &PositionHandler{
		pool: pool,
	}
}

func (h *PositionHandler) GetPosition(c *gin.Context) {
	userID := c.Param("userId")

	var (
		id        string
		userIDDB  string
		symbol    string
		quantity  float64
		entryPrice float64
	)

	err := h.pool.QueryRow(
		c,
		`
		SELECT id, user_id, symbol, quantity, entry_price
		FROM positions
		WHERE user_id = $1
		`,
		userID,
	).Scan(
		&id,
		&userIDDB,
		&symbol,
		&quantity,
		&entryPrice,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Position not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch position",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         id,
		"userId":     userIDDB,
		"symbol":     symbol,
		"quantity":   quantity,
		"entryPrice": entryPrice,
	})
}