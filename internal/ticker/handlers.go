package ticker

import (
	"chaos-go/internal/datasource"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

type Handler struct {
	tickerService TickerService
}

func (handler Handler) AveragePriceHandler(c *gin.Context) {
	timerange := TimeRange{}
	if err := c.ShouldBind(&timerange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	symbol := c.Param("symbol")

	resp, error := handler.tickerService.GetAveragePrice(symbol, timerange.Start, timerange.End)

	if error != nil {
		panic(error)
	}

	c.JSON(http.StatusOK, resp)
}

func (handler Handler) LastPriceHandler(c *gin.Context) {
	symbol := c.Param("symbol")
	resp, error := handler.tickerService.GetLastPrice(symbol)

	if error != nil {
		panic(error)
	}

	c.JSON(http.StatusOK, resp)
}

func (handler Handler) GetPriceByDateHandler(c *gin.Context) {
	date, _ := time.ParseInLocation(time.RFC3339, c.Param("date"), time.UTC)
	accurated := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, time.UTC)
	symbol := c.Param("symbol")
	resp, error := handler.tickerService.GetPriceByDate(symbol, accurated)

	if error != nil {
		if error.Error() == "record not found" {
			println("Get data from provider")
			aggs, err := datasource.GetPriceByDate("BTC", accurated)

			if err != nil {
				panic(err)
			}

			if len(aggs.Entries) > 0 {
				aggregate := &Aggregates{
					Price:  aggs.Entries[0][1],
					Symbol: symbol,
					Date:   accurated,
				}
				SaveOne(aggregate)
				c.JSON(http.StatusOK, aggregate)
				return
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

func (handler Handler) CreateCronJob() {
	cron := cron.New()
	cron.AddFunc("*/1 * * * *", func() {
		symbol := "BTCUSD"
		resp, _ := datasource.GetLatestPrice()

		_, error := handler.tickerService.tickerRepository.GetOrCreateTicker(symbol)

		if error != nil {
			println(error.Error())
			return
		}

		SaveOne(&Aggregates{
			Symbol: symbol,
			Price:  resp.Bpi.USD.Price,
			Date:   time.Now(),
		})
	})

	cron.Start()
}

func NewHandler(tickerService TickerService) *Handler {
	return &Handler{
		tickerService: tickerService,
	}
}
