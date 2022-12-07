package ticker

import (
	"chaos-go/internal/database"
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
	Symbol string `gorm:"not null;index:symbol_date"`
	Price  float64
	Date   time.Time `gorm:"index:symbol_date,sort:desc"`
}

func SaveOne(data interface{}) error {
	db := database.GetDB()
	err := db.Create(data)

	return err.Error
}

func (service TickerService) GetLastPrice(symbol string) (Aggregates, error) {
	var aggregate Aggregates
	db := database.GetDB()
	resp := db.Table("aggregates").Last(&aggregate, "symbol = ?", symbol)

	return aggregate, resp.Error
}

type AvgPrice struct {
	Average float64 `json:"average"`
}

func (service TickerService) GetAveragePrice(symbol string, start time.Time, end time.Time) (AvgPrice, error) {
	var result AvgPrice
	db := database.GetDB()
	resp := db.
		Raw("SELECT AVG(price) average FROM aggregates WHERE symbol = ? AND (date >= ? AND date <= ?)", symbol, start.Format(time.RFC3339), end.Format(time.RFC3339)).
		First(&result)

	return result, resp.Error
}

func (service TickerService) GetPriceByDate(symbol string, date time.Time) (Aggregates, error) {
	var aggregate Aggregates
	db := database.GetDB()
	resp := db.Table("aggregates").First(&aggregate, "symbol = ? AND date = ?", symbol, date)

	return aggregate, resp.Error
}
