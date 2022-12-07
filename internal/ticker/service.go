package ticker

import "time"

type ITickerService interface {
	GetLastPrice(symbol string) (Aggregates, error)
	GetAveragePrice(symbol string, start time.Time, end time.Time) (AvgPrice, error)
	GetPriceByDate(symbol string, date time.Time) (Aggregates, error)
}

type TickerService struct {
	tickerRepository TickerRepository
}
