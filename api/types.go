package main

// struct to represent inventory item
type InventoryItem struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}
