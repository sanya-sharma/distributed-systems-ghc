package handler

import (
	"catalog/models"
	"catalog/service"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// GetCatalog is the api used to get the saree catalog
func GetCatalog(c *gin.Context) {

	//var getCatalogResponse *models.CatalogRequest
	//if err := c.ShouldBindJSON(&getCatalogResponse); err != nil {
	//	log.Printf("Error while parsing order data: %v", err)
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := c.Get("db")

	sareesByCategory, err := service.GetCatalog(db.(*gorm.DB))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting the catalog"})
		return
	}

	result, _ := json.Marshal(sareesByCategory)

	var sarees []*models.Catalog
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
	// Create a context with a timeout of 5 seconds
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, _ := c.Get("db")
	err = service.UpdateCatalogByCategory(db.(*gorm.DB), updateCatalogResponse.CategoryID, updateCatalogResponse.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating the catalog"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, "Updated the catalog successfully")
	return
}
