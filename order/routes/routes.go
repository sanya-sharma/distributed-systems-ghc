package routes

import (
	"github.com/gin-gonic/gin"
	"order/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/place-order", handler.PlaceOrder)
}
