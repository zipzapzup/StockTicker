package api

import (
	"net/http"
	"stockticker/pkg/dates"
	"stockticker/pkg/stock"

	"github.com/gin-gonic/gin"
)

// type album struct {
// 	ID    string `json:"id"`
// 	Title string `json:"title"`
// }

func GetStock(c *gin.Context) {

	TickerInput := "MSFT"
	Apikey := "C227WD9W3LUVKVV9"
	NDAYS := 7

	// Limitations: 500 requests per day
	// Free version: 5 RPM
	// ToDo Implement Rate Limiting

	// var albums = []album{
	// 	{ID: "20", Title: "HelloWorld"},
	// 	{ID: "21", Title: "Hello There"},
	// }
	RespByte := stock.ClientGetStock(TickerInput, Apikey)
	RespStock := stock.ClientConvert(RespByte)

	RespData := stock.ClientFilter(RespStock, TickerInput, dates.ListDates(NDAYS))
	RespData = stock.CalculateAveragePrice(RespData)

	c.IndentedJSON(http.StatusOK, RespData)
}
