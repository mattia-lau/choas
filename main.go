package main

import (
	"chaos-go/internal/database"
	"chaos-go/internal/ticker"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create Database Connection
	db := database.Init("choas.db")
	db.AutoMigrate(&ticker.Ticker{}, &ticker.Aggregates{})

	tickerService := &ticker.TickerService{}
	handlers := ticker.NewHandler(*tickerService)

	// Create Cronjob
	handlers.CreateCronJob()

	// Start WebServer
	r := gin.Default()

	r.GET("/last/ticker/:symbol", handlers.LastPriceHandler)
	r.GET("/average/ticker/:symbol", handlers.AveragePriceHandler)
	r.GET("/:date/ticker/:symbol", handlers.GetPriceByDateHandler)
	r.Run()
}
