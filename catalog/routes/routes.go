package routes

import (
	"catalog/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/get-catalog", handler.GetCatalog)
	r.POST("/update-catalog", handler.UpdateCatalog)

}
