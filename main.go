package main

import (
	"chaos-go/common"
	"chaos-go/datasource"
	"chaos-go/ticker"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func fetchCryptoCron() {
	cron := cron.New()
	cron.AddFunc("*/1 * * * *", func() {
		symbol := "BTCUSD"
		resp, _ := datasource.FetchTicker()

		_, error := ticker.GetOrCreateTicker(symbol)

		if error != nil {
			println(error.Error())
			return
		}

		ticker.SaveOne(&ticker.Aggregates{
			Symbol: symbol,
			Price:  resp.USD.Last,
			Date:   time.Now(),
		})
	})

	cron.Start()
}

func main() {
	_, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create Database Connection
	db := common.Init("choas.db")
	db.AutoMigrate(&ticker.Ticker{}, &ticker.Aggregates{})

	// Create Cronjob
	fetchCryptoCron()

	// Start WebServer
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "On",
		})
	})
	r.GET("/last/ticker/:symbol", ticker.HandleLastPrice)
	r.GET("/average/ticker/:symbol", ticker.HandleAveragePrice)
	r.GET("/:date/ticker/:symbol", ticker.HandleGetPriceByDate)
	r.Run()
}
