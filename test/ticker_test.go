package test

import (
	"chaos-go/internal/database"
	"chaos-go/internal/ticker"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSaveTicker(t *testing.T) {
	database.Init("test.db")
	database.GetDB().AutoMigrate(ticker.Ticker{}, ticker.Aggregates{})
	asserts := assert.New(t)

	now := time.Now()

	data := &ticker.Ticker{
		Symbol: "BNBUSD",
	}
	ticker.SaveOne(data)

	asserts.Equal(data.Symbol, "BNBUSD")
	asserts.GreaterOrEqual(data.CreatedAt.UnixMilli(), now.UnixMilli())
}

func TestSaveAggreate(t *testing.T) {
	database.Init("test.db")
	database.GetDB().AutoMigrate(ticker.Ticker{}, ticker.Aggregates{})
	asserts := assert.New(t)

	now := time.Now()

	data := &ticker.Aggregates{
		Symbol: "BNBUSD",
		Price:  457.7,
		Date:   now,
	}
	ticker.SaveOne(data)

	asserts.Equal(data.Symbol, "BNBUSD")
	asserts.Equal(data.Price, 457.7)
	asserts.Equal(data.Date.UnixMilli(), now.UnixMilli())
}
