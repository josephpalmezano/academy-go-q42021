package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func TestValidateEvenType(t *testing.T) {
	ids := []int{2, 4, 6, 8}
	testService := NewCarService(nil)
	for _, id := range ids {
		testService.WhatType(id)
		assert.Equal(t, "even", "even id.")
	}
}
func TestValidateOddType(t *testing.T) {
	ids := []int{1, 3, 5, 7, 9}
	testService := NewCarService(nil)
	for _, id := range ids {
		testService.WhatType(id)
		assert.Equal(t, "odd", "od id.")
	}
}
