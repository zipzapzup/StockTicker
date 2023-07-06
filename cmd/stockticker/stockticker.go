package main

import (
	"log"

	"stockticker/pkg/api"

	"github.com/gin-gonic/gin"
)

var LocalAddress = "0.0.0.0:8080"

func main() {
	log.Println("Starting Web Server")

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/", api.GetStock)
	router.Run(LocalAddress)
}
