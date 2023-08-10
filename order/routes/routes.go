package routes

import (
	"distributed-systems-ghc/order/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/place-order", handler.PlaceOrder)
}
