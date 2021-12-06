package repository

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"main/model"
	"os"
	"strconv"
)

type CarRepo interface {
	Save(car *model.Car) (*model.Car, error)
	GetAll() ([]model.Car, error)
}
type repo struct{}

func NewCarRepo() CarRepo {
	return &repo{}
}

func listData(data [][]string) []model.Car {
	// convert csv lines to array of structs
	var carList []model.Car
	for i, line := range data {
		if i > 0 { // omit header line
			var rec model.Car
			for j, field := range line {
				if j == 0 {
					var err error
					rec.Id, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				} else if j == 1 {
					var err error
					rec.Year, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
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

func (*repo) GetAll() ([]model.Car, error) {
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

	cars := listData(data)

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
