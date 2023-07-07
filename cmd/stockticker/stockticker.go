package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"stockticker/pkg/dates"
	"stockticker/pkg/stock"

	"github.com/gin-gonic/gin"
)

var LocalAddress = "0.0.0.0:8080"
var GlobalAPIKEY string
var GlobalSYMBOL string
var GlobalNDAYS int

func init() {
	GlobalAPIKEY = os.Getenv("SECRET_APIKEY")
	// Sanitise input
	GlobalAPIKEY = strings.ReplaceAll(GlobalAPIKEY, "\n", "")

	GlobalSYMBOL = os.Getenv("VAR_SYMBOL")
	ndays, _ := strconv.Atoi(os.Getenv("VAR_NDAYS"))
	GlobalNDAYS = ndays
	log.Println(GlobalAPIKEY, GlobalSYMBOL, GlobalNDAYS)
	log.Println("Server initialised: environment variable set")
}

func main() {
	log.Println("Starting Web Server")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", GetStock)
	router.Run(LocalAddress)
}

func GetStock(c *gin.Context) {

	// Limitations: 500 requests per day
	// Free version: 5 RPM
	// ToDo Implement Rate Limiting

	RespByte := stock.ClientGetStock(GlobalSYMBOL, GlobalAPIKEY)
	RespStock := stock.ClientConvert(RespByte)

	RespData := stock.ClientFilter(RespStock, GlobalSYMBOL, dates.ListDates(GlobalNDAYS))
	RespData = stock.CalculateAveragePrice(RespData)

	c.IndentedJSON(http.StatusOK, RespData)
}
