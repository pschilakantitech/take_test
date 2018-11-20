package main

import (
	"encoding/json"
	"io"
)

type RequestParams struct {
	Service string `json:"service,omitempty"`
	Channel string `json:"channel,omitempty"`
	User    struct {
		Seed string `json:"seed,omitempty"`
		Hash string `json:"hash,omitempty"`
	} `json:"user"`
	Coins struct {
		ID        int     `json:"id,omitempty"`
		Name      string  `json:"name,omitempty"`
		CoinCount float64 `json:"coin_count,omitempty"`
	} `json:"coins"`
}

type ResponseQuotes struct {
	Quotes struct {
		USD float64 `json:"USD"`
		INR float64 `json:"INR"`
		BTC float64 `json:"BTC"`
	} `json:"quotes"`
}

func parseRequestBody(body io.ReadCloser) (reqPara *RequestParams, err error) {
	err = json.NewDecoder(body).Decode(&reqPara)
	return
}
