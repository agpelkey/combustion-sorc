package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbtimeout = 5 * time.Second

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Init() (*MongoInstance, error) {
	// set the context
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	// get MongoDB password from env. variable
	//	dbPasswd := os.Getenv("$MONGOPASSWORD")

	// connect to DB with connection URI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://apelkey:CombustionSorc2@cluster0.zpbcywr.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Panic(err)
	}

	// disconnect if there was an error
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Database collection
	db := client.Database("ItemPrices").Collection("prices")

	// initialize MongoInstance
	mg := MongoInstance{
		Client: client,
		DB:     db.Database(),
	}

	return &mg, nil

}
