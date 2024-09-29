package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	mongoURL := "mongodb://localhost:27017"

	fmt.Println("Mongo URL :", mongoURL)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		fmt.Println("Error In Connectin Database : ", err.Error())
		log.Panic(err)
	}

	fmt.Println("Database Connectted Successfully : ")

	return client

}

var Client *mongo.Client = DBInstance()

func OpenCollection(client mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}
