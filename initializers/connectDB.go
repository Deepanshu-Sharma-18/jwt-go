package initializers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func ConnectDB() *mongo.Collection {
	var mongoUri = "mongodb+srv://deepanshursharma:DeepanshuSharma@cluster0.kx1q7pa.mongodb.net/?retryWrites=true&w=majority"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		mongoUri,
	))

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println("There was a problem connecting to your Atlas cluster. Check that the URI includes a valid username and password, and that your IP address has been added to the access list. Error: ")
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	var dbName = "myDatabase"
	var collectionName = "users"
	clc := client.Database(dbName).Collection(collectionName)

	clc.InsertOne(ctx, bson.D{primitive.E{Key: "name", Value: "Deepanshu"}})

	Collection = clc

	return clc
}
