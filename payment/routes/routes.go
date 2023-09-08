package routes

import (
	"github.com/gin-gonic/gin"
	"payment/handler"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/get-available-payment-methods", handler.GetAvailablePaymentMethods)
	r.POST("/initiate-payment", handler.InitiatePayment)
	r.POST("/rollback-payment", handler.RollbackPayment)
}
