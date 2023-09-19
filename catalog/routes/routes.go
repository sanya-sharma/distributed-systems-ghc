package routes

import (
	"catalog/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/get-catalog", handler.GetCatalog)
	r.GET("/get-catalog-by-productid", handler.GetCatalogByProductID)
	r.POST("/update-catalog", handler.UpdateCatalog)

}
