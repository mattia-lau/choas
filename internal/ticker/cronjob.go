package ticker

import (
	"chaos-go/internal/datasource"
	"time"

	"github.com/robfig/cron/v3"
)

func CreateCronJob(repository TickerRepository) {
	cron := cron.New()
	cron.AddFunc("*/1 * * * *", func() {
		symbol := "BTCUSD"
		resp, _ := datasource.GetLatestPrice()

		_, error := repository.GetOrCreateTicker(symbol)

		if error != nil {
			println(error.Error())
			return
		}

		repository.SaveOne(&Aggregates{
			Symbol: symbol,
			Price:  resp.Bpi.USD.Price,
			Date:   time.Now(),
		})
	})

	cron.Start()
}
