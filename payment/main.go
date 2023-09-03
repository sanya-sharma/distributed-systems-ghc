package main

import (
	"distributed-systems-ghc/payment/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Next()
	})
	routes.SetupRoutes(router)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

