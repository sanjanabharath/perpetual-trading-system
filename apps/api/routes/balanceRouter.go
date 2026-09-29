package routes

import gin "github.com/gin-gonic/gin"

var router = gin.Default()

func init() {
	router.GET("/balance", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"balance": 1000.00,
		})
	})
}