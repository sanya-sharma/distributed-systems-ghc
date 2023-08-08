package controller

import (
	"context"
	servercontroller "distributed-systems-ghc/distributed-systems-ghc/server/controller"
	"net/http"
	"time"
)

// GetCatalog is the api used to get the saree catalog
func GetCatalog(response http.ResponseWriter, request *http.Request) (err error) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	response.Header().Add("Content-Type", "application/json")

	err = servercontroller.GetCatalog(response, request)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	/*
		todo: client side code might include using the business results somehow
	*/

	return nil
}
