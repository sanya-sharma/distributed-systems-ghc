package routes

import (
	"distributed-systems-ghc/payment/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/get-payment-methods", handler.GetAvailablePaymentMethods)

}
