package ticker

import (
	"chaos-go/datasource"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleAveragePrice(c *gin.Context) {
	timerange := TimeRange{}
	if err := c.ShouldBind(&timerange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": fmt.Sprintf("%v", err)})
		return
	}

	symbol := c.Param("symbol")

	resp, error := GetAveragePrice(symbol, timerange.Start, timerange.End)

	if error != nil {
		panic(error)
	}

	c.JSON(http.StatusOK, resp)
}

func HandleLastPrice(c *gin.Context) {
	symbol := c.Param("symbol")
	resp, error := GetLastPrice(symbol)

	if error != nil {
		panic(error)
	}

	c.JSON(http.StatusOK, resp)
}

func HandleGetPriceByDate(c *gin.Context) {
	date, _ := time.ParseInLocation(time.RFC3339, c.Param("date"), time.UTC)
	accurated := time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), 0, 0, time.UTC)
	symbol := c.Param("symbol")
	resp, error := GetPriceByDate(symbol, accurated)

	if error != nil {
		if error.Error() == "record not found" {
			println("Get data from provider")
			aggs, err := datasource.FetchPolygonTicker(symbol, accurated)

			if err != nil {
				panic(err)
			}

			if aggs.Count > 0 {
				aggregate := &Aggregates{
					Price:  aggs.Results[0].Close,
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
