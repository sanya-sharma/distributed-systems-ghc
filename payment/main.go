package main

import (
	"github.com/gin-gonic/gin"
	"payment/routes"
)

func main() {

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Next()
	})
	routes.SetupRoutes(router)
	router.Run(":8082") // listen and serve on 0.0.0.0:8082

}
