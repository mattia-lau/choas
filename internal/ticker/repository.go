package ticker

import "chaos-go/internal/database"

// go:generate mockery --name ITickerRepository --filename ticker_repository_mock.go
type ITickerRepository interface {
	SaveOne(data Ticker) error
	GetOrCreateTicker(symbol string) (Ticker, error)
}

type TickerRepository struct {
}

func (repo TickerRepository) SaveOne(data Ticker) error {
	db := database.GetDB()
	err := db.Create(data)

	return err.Error
}

func (repo TickerRepository) GetOrCreateTicker(symbol string) (Ticker, error) {
	var ticker Ticker
	db := database.GetDB()
	resp := db.Table("tickers").FirstOrCreate(&ticker, Ticker{Symbol: symbol})

	return ticker, resp.Error
}
