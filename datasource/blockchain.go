package datasource

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type Response struct {
	USD struct {
		Last float64 `json:"last"`
	} `json:"USD"`
}

func FetchTicker() (*Response, error) {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://blockchain.info/ticker")

	response := &Response{}
	json.Unmarshal(resp.Body(), response)

	return response, err
}
