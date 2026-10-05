package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DepthHandler struct {
	pool *pgxpool.Pool
}

func NewDepthHandler(pool *pgxpool.Pool) *DepthHandler {
	return &DepthHandler{
		pool: pool,
	}
}

type Depth struct {
	ID string `json:"id"`
	// Add your actual depth fields here
}

func (h *DepthHandler) GetDepth(c *gin.Context) {
	var depth Depth

	err := h.pool.QueryRow(
		c,
		`
		SELECT id
		FROM depth
		WHERE id = $1
		`,
		"BTCUSDT",
	).Scan(
		&depth.ID,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Depth not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch depth",
		})
		return
	}

	c.JSON(http.StatusOK, depth)
}