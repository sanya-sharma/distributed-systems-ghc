package routes

import (
	"distributed-systems-ghc/payment/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/get-available-payment-methods", handler.GetAvailablePaymentMethods)
	r.POST("/initiate-payment", handler.InitiatePayment)
	r.POST("/rollback-payment", handler.RollbackPayment)
}
