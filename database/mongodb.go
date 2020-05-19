package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	err    error
)

// ConnectDB : Takes in a mongodb connection string and connects to the database
func ConnectDB(uri string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		panic("Cannot establish database client")
	}
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
		panic("Database connection error")
	}
	fmt.Println("MongoDB connection successful")
}

// GetDB : Returns the database client
func GetDB() *mongo.Client {
	return client
}
