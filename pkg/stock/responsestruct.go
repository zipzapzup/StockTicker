package stock

type StockResponse struct {
	StockName    string
	AveragePrice float64
	StockPrice   map[string]string
}
