package servercontroller

import (
	"context"
	"distributed-systems-ghc/models"
	"distributed-systems-ghc/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// GetCatalog is the api used to get the saree catalog
func GetCatalog(response http.ResponseWriter, request *http.Request) (err error) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	response.Header().Add("Content-Type", "application/json")

	queryParams := request.URL.Query()
	// Extract the "categoryID" query parameter
	categoryIDParam := queryParams.Get("categoryID")
	//Extract the "colorID" query param
	colorIDParam := queryParams.Get("color")

	sareesByCategory, err := service.GetCatalogByCategory(categoryIDParam, colorIDParam)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	fmt.Print(response)

	result, _ := json.Marshal(sareesByCategory)

	var sarees []*models.Product
	if err = json.Unmarshal(result, &sarees); err != nil {
		log.Fatal(err)
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	err = json.NewEncoder(response).Encode(sarees)
	if err != nil {
		//todo: log here
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return err
	}
	return nil
}

// GetCatalog is the api used to get the saree catalog
func PlaceOrder(response http.ResponseWriter, request *http.Request) (err error) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	response.Header().Add("Content-Type", "application/json")

	queryParams := request.URL.Query()
	// Extract the "categoryID" query parameter
	categoryIDParam := queryParams.Get("categoryID")
	//Extract the "colorID" query param
	colorIDParam := queryParams.Get("color")

	sareesByCategory, err := service.GetCatalogByCategory(categoryIDParam, colorIDParam)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}

	fmt.Print(response)

	result, _ := json.Marshal(sareesByCategory)

	var sarees []*models.Product
	if err = json.Unmarshal(result, &sarees); err != nil {
		log.Fatal(err)
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	err = json.NewEncoder(response).Encode(sarees)
	if err != nil {
		//todo: log here
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return err
	}
	return nil
}
