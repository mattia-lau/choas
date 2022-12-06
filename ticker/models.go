package ticker

import (
	"chaos-go/common"
	"time"
)

type Ticker struct {
	Symbol    string       `gorm:"primaryKey;not null"`
	Prices    []Aggregates `gorm:"ForeignKey:Symbol"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Aggregates struct {
	ID     uint   `gorm:"primarykey"`
	Symbol string `gorm:"not null"`
	Price  float64
	Date   time.Time
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	print(db)
	err := db.Create(data)

	return err.Error
}

func GetOrCreateTicker(symbol string) (Ticker, error) {
	var ticker Ticker
	db := common.GetDB()
	resp := db.Table("tickers").FirstOrCreate(&ticker, Ticker{Symbol: symbol})

	return ticker, resp.Error
}

func GetLastPrice(symbol string) (Aggregates, error) {
	var aggregate Aggregates
	db := common.GetDB()
	resp := db.Table("aggregates").Last(&aggregate, "symbol = ?", symbol)

	return aggregate, resp.Error
}

type AvgPrice struct {
	Average float64 `json:"average"`
}

func GetAveragePrice(symbol string, start time.Time, end time.Time) (AvgPrice, error) {
	var result AvgPrice
	db := common.GetDB()
	resp := db.
		Raw("SELECT AVG(price) average FROM aggregates WHERE symbol = ? AND (date >= ? AND date <= ?)", symbol, start.Format(time.RFC3339), end.Format(time.RFC3339)).
		First(&result)

	return result, resp.Error
}

func GetPriceByDate(symbol string, date time.Time) (Aggregates, error) {
	var aggregate Aggregates
	db := common.GetDB()
	resp := db.Table("aggregates").First(&aggregate, "symbol = ? AND date = ?", symbol, date)

	return aggregate, resp.Error
}
