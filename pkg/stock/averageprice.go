// package stock specifically used to calculate the average price
package stock

import (
	"fmt"
	"math"
	"strconv"
)

func CalculateAveragePrice(Stock StockResponse) StockResponse {
	counter := 0.0
	total := 0.0
	for _, v := range Stock.StockPrice {
		counter += 1.0
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		}
		total += num
	}

	result := total / float64(counter)

	// Round to 2 decimal places
	Roundresult := math.Round(result*100) / 100
	Stock.AveragePrice = float64(Roundresult)

	return Stock
}
