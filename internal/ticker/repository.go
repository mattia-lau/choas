package ticker

import (
	"chaos-go/internal/database"
	"time"

	"gorm.io/gorm"
)

type TickerRepository interface {
	SaveOne(data interface{}) error
	GetLastPrice(symbol string) (Aggregates, error)
	GetAveragePrice(symbol string, start time.Time, end time.Time) (AvgPrice, error)
	GetPriceByDate(symbol string, date time.Time) (Aggregates, error)
	GetOrCreateTicker(symbol string) (Ticker, error)
	GetTicker(symbol string) (Ticker, error)
}

type tickerDB struct {
	db *gorm.DB
}

func (a tickerDB) SaveOne(data interface{}) error {
	db := database.GetDB()
	err := db.Create(data)

	return err.Error
}

func (a tickerDB) GetLastPrice(symbol string) (Aggregates, error) {
	var aggregate Aggregates
	db := database.GetDB()
	resp := db.Table("aggregates").Last(&aggregate, "symbol = ?", symbol)

	return aggregate, resp.Error
}

func (a tickerDB) GetAveragePrice(symbol string, start time.Time, end time.Time) (AvgPrice, error) {
	var result AvgPrice
	db := database.GetDB()
	resp := db.
		Raw("SELECT AVG(price) average FROM aggregates WHERE symbol = ? AND (date >= ? AND date <= ?)", symbol, start.Format(time.RFC3339), end.Format(time.RFC3339)).
		First(&result)

	return result, resp.Error
}

func (a tickerDB) GetPriceByDate(symbol string, date time.Time) (Aggregates, error) {
	var aggregate Aggregates
	db := database.GetDB()
	resp := db.Table("aggregates").First(&aggregate, "symbol = ? AND date = ?", symbol, date)

	return aggregate, resp.Error
}

func (a tickerDB) GetOrCreateTicker(symbol string) (Ticker, error) {
	var ticker Ticker
	db := database.GetDB()
	resp := db.Table("tickers").FirstOrCreate(&ticker, Ticker{Symbol: symbol})

	return ticker, resp.Error
}

func (a tickerDB) GetTicker(symbol string) (Ticker, error) {
	var ticker Ticker
	db := database.GetDB()
	resp := db.Table("tickers").First(&ticker, Ticker{Symbol: symbol})

	return ticker, resp.Error
}

func NewTickerRepository(db *gorm.DB) TickerRepository {
	return &tickerDB{
		db: db,
	}
}
