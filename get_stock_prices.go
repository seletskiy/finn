package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	karma "github.com/reconquest/karma-go"
)

const (
	urlBatchStockQuotes = "/query?function=BATCH_STOCK_QUOTES&symbols=%s&apikey=%s"
)

func getStockPrices(
	apiKey string, symbols []string,
) (map[string]float64, error) {
	response, err := http.Get(
		apiHost + fmt.Sprintf(
			urlBatchStockQuotes,
			strings.Join(symbols, ","),
			apiKey,
		),
	)
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to get batch stock quotes",
		)
	}

	defer response.Body.Close()

	quotes := BatchStockQuotes{}

	err = json.NewDecoder(response.Body).Decode(&quotes)
	if err != nil {
		return nil, karma.Format(
			err,
			"unable to decode batch stock quotes",
		)
	}

	prices := map[string]float64{}

	for _, quote := range quotes.StockQuotes {
		price, err := strconv.ParseFloat(quote.Price, 0)
		if err != nil {
			return nil, karma.Describe("value", quote.Price).Format(
				err,
				"unable to parse price as float64",
			)
		}

		prices[quote.Symbol] = price
	}

	return prices, nil
}
