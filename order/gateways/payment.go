package gateways 

import (
    "bytes"
    "encoding/json"
    "net/http"
	"order/models"
	"order/config"
	"log"
)

func InitiatePayment(order models.Order) error {
    payment := models.Payment{
        OrderID:       order.ID,
        PaymentMethod: "Paypal", // TODO: Specify the payment method
        Amount:        calculatePaymentAmount(order), // TODO: Calculate the payment amount
        Status:        "Pending",
    }

    requestBody, err := json.Marshal(payment)
    if err != nil {
        return err
    }

    paymentServiceURL, err := config.ReadServiceConfig("payment")
    if err != nil {
        return err
    }
	initiatePaymentRoute, err:= config.ReadAPIConfig("initiate-payment")
    if err != nil {
        return err
    }
	initiatePaymentURL := paymentServiceURL + initiatePaymentRoute

    _, err = http.Post(initiatePaymentURL, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return err
    }

    return nil
}

func RollbackPayment(order models.Order) error {
    payment := models.Payment{
        OrderID:       order.ID,
        PaymentMethod: "Paypal", // TODO: Specify the payment method
        Amount:        calculatePaymentAmount(order), // TODO: Calculate the payment amount
        Status:        "Pending",
    }

    requestBody, err := json.Marshal(payment)
    if err != nil {
        return err
    }

    paymentServiceURL, err := config.ReadServiceConfig("payment")
    if err != nil {
        return err
    }
	rollbackPaymentRoute, err:= config.ReadAPIConfig("rollback-payment")
    if err != nil {
        return err
    }
	rollbackPaymentURL := paymentServiceURL + rollbackPaymentRoute

    _, err = http.Post(rollbackPaymentURL, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return err
    }

    return nil
}


func calculatePaymentAmount(order models.Order) float64 {
	return 10.0
}