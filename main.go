package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	apiKey = "UYWSV9DWV1GCKGBB"
)

func main() {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("unable to read stdin: %s", err)
	}

	symbols := strings.Split(
		strings.TrimSpace(string(stdin)),
		"\n",
	)

	prices, err := getStockPrices(apiKey, symbols)
	if err != nil {
		log.Fatalf("unable to get stock prices: %s", err)
	}

	for _, symbol := range symbols {
		fmt.Printf("%f\n", prices[symbol])
	}
}
