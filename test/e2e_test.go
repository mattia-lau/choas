package test

import (
	"chaos-go/internal/config"
	"chaos-go/internal/database"
	"chaos-go/internal/logging"
	"chaos-go/internal/ticker"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TickerSuite struct {
	suite.Suite
	handler    *ticker.Handler
	repository ticker.TickerRepository
}

func (s *TickerSuite) SetupTest() {
	cfg, _ := config.Load("./test.yml")
	logger := logging.SetupLogger(cfg)
	db := database.Init(cfg, logger)
	s.repository = ticker.NewTickerRepository(db)
	s.handler = ticker.NewHandler(s.repository)
}

func TestConnectingDatabase(t *testing.T) {
	dbPath := "test.db"
	asserts := assert.New(t)

	cfg, _ := config.Load("./test.yml")

	logger := logging.SetupLogger(cfg)
	db, _ := database.Init(cfg, logger).DB()
	// Test create & close DB
	_, err := os.Stat(dbPath)
	asserts.NoError(err, "Db should exist")
	asserts.NoError(db.Ping(), "Db should be able to ping")

	// Test get a connecting from connection pools
	connection := database.GetDB()
	asserts.NoError(connection.Error, "Db should be able to ping")
	defer db.Close()
}

func (s *TickerSuite) TestSaveTicker(t *testing.T) {
	database.GetDB().AutoMigrate(ticker.Ticker{}, ticker.Aggregates{})
	asserts := assert.New(t)

	now := time.Now()

	data := ticker.Ticker{
		Symbol: "BNBUSD",
	}
	s.repository.SaveOne(data)

	asserts.Equal(data.Symbol, "BNBUSD")
	asserts.GreaterOrEqual(data.CreatedAt.UnixMilli(), now.UnixMilli())
}

func (s *TickerSuite) TestSaveAggreate(t *testing.T) {
	database.GetDB().AutoMigrate(ticker.Ticker{}, ticker.Aggregates{})
	asserts := assert.New(t)

	now := time.Now()

	data := &ticker.Aggregates{
		Symbol: "BNBUSD",
		Price:  457.7,
		Date:   now,
	}
	s.repository.SaveOne(data)

	asserts.Equal(data.Symbol, "BNBUSD")
	asserts.Equal(data.Price, 457.7)
	asserts.Equal(data.Date.UnixMilli(), now.UnixMilli())
}
