package main

import (
	"log"
	"stockticker/pkg/api"

	"github.com/gin-gonic/gin"
)

var LocalAddress = "localhost:8080"

func main() {
	log.Println("Starting Web Server")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", api.GetStock)
	router.Run(LocalAddress)
}
