package handler

import (
	"distributed-systems-ghc/catalog/models"
	"distributed-systems-ghc/catalog/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetCatalog is the api used to get the saree catalog
func GetCatalog(c *gin.Context) {

	requestBody, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	var getCatalogResponse *models.CatalogResponse
	err = json.Unmarshal(requestBody, &getCatalogResponse)

	sareesByCategory, err := service.GetCatalogByCategory(getCatalogResponse.CategoryID, getCatalogResponse.ColorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting the catalog"})
		return
	}

	result, _ := json.Marshal(sareesByCategory)

	var sarees []*models.Product
	if err = json.Unmarshal(result, &sarees); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, sareesByCategory)
	c.Header("Content-Type", "application/json")

	return
}

// UpdateCatalog is the api used to update the saree catalog
func UpdateCatalog(c *gin.Context) {

	requestBody, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	var updateCatalogResponse *models.CatalogResponse
	err = json.Unmarshal(requestBody, &updateCatalogResponse)

	sareesByCategory, err := service.UpdateCatalogByCategory(updateCatalogResponse.CategoryID, updateCatalogResponse.ColorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating the catalog"})
		return
	}

	result, _ := json.Marshal(sareesByCategory)

	var sarees []*models.Product
	if err = json.Unmarshal(result, &sarees); err != nil {
		log.Fatal(err)
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, "Updated the catalog successfully")
	return
}
