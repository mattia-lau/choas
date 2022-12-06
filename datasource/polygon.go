package datasource

import (
	"context"
	"fmt"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func FetchPolygonTicker(symbol string, date time.Time) (*models.GetAggsResponse, error) {
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))
	resp, err := c.GetAggs(context.Background(), &models.GetAggsParams{
		From:       models.Millis(date),
		To:         models.Millis(date),
		Multiplier: 1,
		Timespan:   "minute",
		Ticker:     fmt.Sprintf("X:%s", symbol),
	})

	return resp, err
}
