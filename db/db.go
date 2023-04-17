package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	UsersCollection = "users"
)

var mongoClient *mongo.Client

func SetupMongoDBClient(ctx context.Context) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:2vitWGctqOohMdEn@cluster0.5bbukuf.mongodb.net/?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client
}

func Collection(collection string) (*mongo.Collection, error) {
	// Set a default database for the client
	// Return a collection object bound to the default database
	defaultDatabase := "myData"
	return mongoClient.Database(defaultDatabase).Collection(collection), nil
}
