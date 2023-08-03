package controller

import (
	"context"
	"distributed-systems-ghc/distributed-systems-ghc/models"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// GetCatalog is the api used to get the saree catalog
func GetCatalog(response http.ResponseWriter, request *http.Request) (err error) {
	var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	response.Header().Add("Content-Type", "application/json")
	var result []byte
	//	todo: either a direct db call or a service layer based on what and how much should the user see, finalise arguments
	/*
		result, err := sareeCollection.Find()
		fmt.Print(result)
	*/

	defer cancel()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
	}
	var sarees []*models.Product
	if err = json.Unmarshal(result, sarees); err != nil {
		log.Fatal(err)
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	json.NewEncoder(response).Encode(sarees)
	return nil
}
