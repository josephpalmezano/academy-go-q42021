package service

import (
	"main/repository"
)

type CarService interface {
	WhatType(id int) (isType string)
}

type service struct{}

var (
	repo repository.CarRepo
)

func NewCarService(repository repository.CarRepo) *service {
	repo = repository
	return &service{}
}

func WhatType(id int) (isType string) {
	if id%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
