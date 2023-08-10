package handler

import (
	"distributed-systems-ghc/catalog/models"
	"distributed-systems-ghc/catalog/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetCatalog is the api used to get the saree catalog
func GetCatalog(c *gin.Context) {

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
