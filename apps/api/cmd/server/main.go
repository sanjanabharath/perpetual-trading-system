package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")

	var db *sql.DB

	if databaseURL != "" {
		var err error

		db, err = sql.Open("pgx", databaseURL)
		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		log.Println("Connected to PostgreSQL")
	}

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		response := gin.H{
			"status":  "ok",
			"service": "perpetual-trading-api",
		}

		if db != nil {
			response["database"] = "connected"
		}

		c.JSON(http.StatusOK, response)
	})

	log.Printf(
		"Perpetual Trading API running on http://localhost:%s",
		port,
	)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
