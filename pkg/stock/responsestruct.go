// package stock which creates the StockResponse structure
// StockResponse will be returned to the handler
package stock

type StockResponse struct {
	StockName    string
	AveragePrice float64
	StockPrice   map[string]string
}
