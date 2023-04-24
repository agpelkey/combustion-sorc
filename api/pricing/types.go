package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Price int                `json:"price" bson:"price,omitempty"`
}
