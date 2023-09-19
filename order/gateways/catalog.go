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

func GetCatalog(inventory *models.Catalog) (catalog *models.Catalog, err error) {
	catalogInventory, _ := json.Marshal(inventory)

	requestBody, err := json.Marshal(catalogInventory)
	if err != nil {
		return nil, err
	}

	catalogServiceURL, err := config.ReadServiceConfig("catalog")
	if err != nil {
		return nil, err
	}
	getCatalogRoute, err := config.ReadAPIConfig("get-catalog-by-productid")
	if err != nil {
		return nil, err
	}
	getCatalogURL := catalogServiceURL + getCatalogRoute

	resp, err := http.Post(getCatalogURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			return nil, err
		}
		return nil, errors.New(string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	var sarees *models.Catalog
	err = json.Unmarshal(body, &sarees)
	if err != nil {
		log.Println("Error unmarshalling response body:", err)
		return nil, err
	}

	return sarees, nil
}
