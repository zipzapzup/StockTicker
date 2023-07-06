package stock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StockAlpha struct {
	TimeSeries DateSeries `json:"Time Series (Daily)"`
}

type DateSeries map[string]StockData

type StockData struct {
	Open   string `json:"1. open"`
	Close  string `json:"4. close"`
	Volume string `json:"6. volume"`
}

func ClientGetStock(SYMBOL string, API string) []byte {

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY_ADJUSTED&symbol=%s&outputsize=full&apikey=%s&datatype=json", SYMBOL, API)

	client := &http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return body

}

func ClientConvert(RawBody []byte) StockAlpha {
	// Convert Byte to StockAlpha struct
	// Unmarshal
	NewStock := StockAlpha{}
	err := json.Unmarshal(RawBody, &NewStock)
	if err != nil {
		log.Println(err)
	}

	return NewStock
}

func ClientFilter(Alpha StockAlpha, Name string, Dates []string) StockResponse {
	// Filters Alpha stock with NDAYS dates
	// Two options exists here
	// To remove the data loaded or to create a new map
	// Go with the later

	// One assumption, NDAYS is not based on the sharemarket calendar days

	var Data StockResponse
	Data.StockName = Name
	Data.StockPrice = make(map[string]string)

	// Loop through the dates and create map if found
	for _, v := range Dates {
		_, exists := Alpha.TimeSeries[v]
		if exists {
			Data.StockPrice[v] = Alpha.TimeSeries[v].Close
		}
	}
	return Data

}
