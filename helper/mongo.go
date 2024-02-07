package main

import (
	"context"
	"fmt"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Create MongoDB client
func makeMongoClient() (*mongo.Client, error) {


	// change the below line to your MongoDB connection string with username and password of kubernetes

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping MongoDB to verify connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Insert documents into MongoDB collection
func Insert() error {
	client, err := makeMongoClient()
	if err != nil {
		return err
	}

	// Access database and collection
	db := client.Database("mydb")
	collection := db.Collection("test")

	// Insert multiple documents
	docs := make([]interface{}, 100)
	for i := range docs {
		docs[i] = bson.M{"state": "QUEUED"}
	}
	_, err = collection.InsertMany(context.Background(), docs)
	return err
	
}

// Delete all documents from MongoDB collection
func Delete() error {
	client, err := makeMongoClient()
	if err != nil {
		return err
	}

	// Access database and collection
	db := client.Database("mydb")
	collection := db.Collection("test")

	// Delete all documents
	_, err = collection.DeleteMany(context.Background(), bson.M{})
	return err

}
