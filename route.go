package main

import (
	"encoding/json"
	"main/repository"
	"net/http"
)

var (
	repo repository.CarRepo = repository.NewCarRepo()
)

func GetCars(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	cars, err := repo.GetAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the cars"}`))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(cars)

}
