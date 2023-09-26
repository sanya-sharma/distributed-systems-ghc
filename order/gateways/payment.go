package gateways

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"order/config"
	"order/models"
)

var (
	paymentStatusPending = "Pending"
)

func InitiatePayment(order models.Order) error {
	payment := models.Payment{
		OrderID: order.OrderID,
		Amount:  100.0,
		Status:  paymentStatusPending,
	}

	requestBody, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	paymentServiceURL, err := config.ReadServiceConfig("payment")
	if err != nil {
		return err
	}
	initiatePaymentRoute, err := config.ReadAPIConfig("initiate-payment")
	if err != nil {
		return err
	}
	initiatePaymentURL := paymentServiceURL + initiatePaymentRoute

	resp, err := http.Post(initiatePaymentURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			return err
		}
		return errors.New(string(body))
	}

	return nil
}

func RollbackPayment(order models.Order) error {
	payment := models.Payment{
		OrderID: order.OrderID,
		Amount:  100.00,
		Status:  paymentStatusPending,
	}

	requestBody, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	paymentServiceURL, err := config.ReadServiceConfig("payment")
	if err != nil {
		return err
	}
	rollbackPaymentRoute, err := config.ReadAPIConfig("rollback-payment")
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
