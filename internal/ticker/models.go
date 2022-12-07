package ticker

import "time"

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
