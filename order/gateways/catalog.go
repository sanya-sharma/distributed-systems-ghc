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

func UpdateCatalog(catalog *models.Catalog) (err error) {

	requestBody, err := json.Marshal(catalog)
	if err != nil {
		return err
	}

	catalogServiceURL, err := config.ReadServiceConfig("catalog")
	if err != nil {
		return err
	}
	updateCatalogRoute, err := config.ReadAPIConfig("update-catalog")
	if err != nil {
		return err
	}
	updateCatalogURL := catalogServiceURL + updateCatalogRoute

	resp, err := http.Post(updateCatalogURL, "application/json", bytes.NewBuffer(requestBody))
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
