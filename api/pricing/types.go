package main

type Item struct {
	ID    int `json:"id,omitempty" bson:"_id,omitempty"`
	Price int `json:"price" bson:"price"`
}
