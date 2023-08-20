package routes

import (
	"distributed-systems-ghc/catalog/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/get-catalog", handler.GetCatalog)

}
