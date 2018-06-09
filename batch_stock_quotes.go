package main

type BatchStockQuotes struct {
	StockQuotes []struct {
		Symbol    string    `json:"1. symbol"`
		Price     string    `json:"2. price"`
		Volume    string    `json:"3. volume"`
		Timestamp Timestamp `json:"4. timestamp"`
	} `json:"Stock Quotes"`
}
