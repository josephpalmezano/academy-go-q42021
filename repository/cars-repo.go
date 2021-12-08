package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"main/model"
	"os"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

type CarRepo interface {
	Save(car *model.Car) (*model.Car, error)
	GetAll(reqType string) ([]model.Car, error)
}
type repo struct{}

func NewCarRepo() CarRepo {
	return &repo{}
}

func WhatType(id int) (isType string) {
	if id%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func listData(data [][]string, reqType string) []model.Car {
	var carList []model.Car
	for i, line := range data {
		id, _ := strconv.Atoi(line[0])
		if i > 0 && (WhatType(id) == reqType || reqType == "") {
			var rec model.Car
			for j, field := range line {
				if j == 0 {
					rec.Id, _ = strconv.Atoi(field)
				} else if j == 1 {
					rec.Year, _ = strconv.Atoi(field)
				} else if j == 2 {
					rec.Brand = field
				} else if j == 3 {
					rec.Model = field
				} else if j == 4 {
					rec.Color = field
				}
			}
			carList = append(carList, rec)
		}
	}
	return carList
}

func (*repo) GetAll(reqType string) ([]model.Car, error) {
	f, err := os.Open("data/cars.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var cars []model.Car

	wg.Add(1)

	go func([]model.Car) {
		cars = listData(data, reqType)
		wg.Done()
	}(cars)

	wg.Wait()

	jsonData, err := json.MarshalIndent(cars, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	return cars, err
}

func (*repo) Save(car *model.Car) (*model.Car, error) {
	f, err := os.OpenFile("data/cars.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var row [][]string
	row = append(row, []string{strconv.Itoa(car.Id), strconv.Itoa(car.Year), car.Brand, car.Model, car.Color})

	w := csv.NewWriter(f)
	w.WriteAll(row)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Appending succed")

	return car, nil
}
