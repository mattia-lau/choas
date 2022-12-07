package datasource

import (
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
)

type coindeskCurrentPrice struct {
	Bpi struct {
		USD struct {
			Price float64 `json:"rate_float"`
		} `json:"USD"`
	} `json:"bpi"`
}

func GetLatestPrice() (*coindeskCurrentPrice, error) {
	//
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://api.coindesk.com/v1/bpi/currentprice.json")

	response := &coindeskCurrentPrice{}
	json.Unmarshal(resp.Body(), response)

	return response, err
}

type requestQuery struct {
	EndDate   string `json:"end_date"`
	ISO       string `json:"iso"`
	OHLC      bool   `json:"ohlc"`
	StartDate string `json:"start_date"`
}

type chartApiResponse struct {
	Iso            string      `json:"iso"`
	Name           string      `json:"name"`
	Slug           string      `json:"slug"`
	IngestionStart string      `json:"ingestionStart"`
	Interval       string      `json:"interval"`
	Src            string      `json:"src"`
	Entries        [][]float64 `json:"entries"`
	ID             string      `json:"_id"`
}

func GetPriceByDate(symbol string, date time.Time) (*chartApiResponse, error) {
	client := resty.New()

	query := &requestQuery{
		EndDate:   "2022-12-06T14:09",
		ISO:       symbol,
		OHLC:      false,
		StartDate: "2022-12-06T14:08",
	}

	jsonString, err := json.Marshal(query)

	resp, err := client.R().
		EnableTrace().
		SetQueryParam("query", string(jsonString)).
		Get("https://www.coindesk.com/pf/api/v3/content/fetch/chart-api")

	response := &chartApiResponse{}
	json.Unmarshal(resp.Body(), response)

	return response, err
}
