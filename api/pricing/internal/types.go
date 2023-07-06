package data

import (
	"errors"
)

type Item struct {
	ID    int `json:"id"`
	Name  int `json:"name"`
	Price int `json:"price"`
}

var (
	ErrRecordNotFound = errors.New("record not found")
)
