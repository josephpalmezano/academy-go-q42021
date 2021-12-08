package main

import (
	"encoding/json"
	"main/model"
	"main/repository"
	"math/rand"
	"net/http"
)

var (
	repo repository.CarRepo = repository.NewCarRepo()
)

func GetCars(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	reqType := req.URL.Query().Get("type")
	cars, err := repo.GetAll(reqType)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the cars"}`))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(cars)
}

func AddCar(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var car model.Car
	err := json.NewDecoder(request.Body).Decode(&car)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}
	car.Id = rand.Int()
	repo.Save(&car)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(car)
}
