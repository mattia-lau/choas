package ticker

import (
	"chaos-go/internal/datasource"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	repository TickerRepository
}

func (handler Handler) AveragePriceHandler(c *gin.Context) {
	timerange := TimeRange{}
	if err := c.ShouldBind(&timerange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	symbol := c.Param("symbol")
	if _, err := handler.repository.GetTicker(symbol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s not found", symbol)})
		return
	}

	resp, err := handler.repository.GetAveragePrice(symbol, timerange.Start, timerange.End)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (handler Handler) LastPriceHandler(c *gin.Context) {
	symbol := c.Param("symbol")
	if _, err := handler.repository.GetTicker(symbol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s not found", symbol)})
		return
	}

	resp, err := handler.repository.GetLastPrice(symbol)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (handler Handler) GetPriceByDateHandler(c *gin.Context) {
	validator := DateValidator{}
	if err := c.ShouldBindUri(&validator); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	date := validator.Date
	symbol := validator.Symbol

	if _, err := handler.repository.GetTicker(symbol); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%s not found", symbol)})
		return
	}

	accurated := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, time.UTC)
	resp, error := handler.repository.GetPriceByDate(symbol, accurated)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		println("Get data from provider", accurated.Format(time.RFC3339))
		aggs, err := datasource.GetPriceByDate("BTC", accurated)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
			return
		}

		if len(aggs.Entries) > 0 {
			aggregate := &Aggregates{
				Price:  aggs.Entries[0][1],
				Symbol: symbol,
				Date:   accurated,
			}
			handler.repository.SaveOne(aggregate)
			c.JSON(http.StatusOK, aggregate)
			return
		}
	}

	c.JSON(http.StatusOK, resp)
}

func NewHandler(repository TickerRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

func Route(r *gin.Engine, h *Handler) {
	r.GET("/last/ticker/:symbol", h.LastPriceHandler)
	r.GET("/average/ticker/:symbol", h.AveragePriceHandler)
	r.GET("/:date/ticker/:symbol", h.GetPriceByDateHandler)
}
