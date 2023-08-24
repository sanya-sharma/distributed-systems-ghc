package main

import (
	"distributed-systems-ghc/order/routes"
	"github.com/cavdy-play/go_db/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Connect()

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
