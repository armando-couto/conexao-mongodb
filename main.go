package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// ConnectToDB starts a new database connection and returns a reference to it
func ConnectToDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	options := options.Client().ApplyURI("mongodb://localhost:32768")
	options.SetMaxPoolSize(1000)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	return client.Database("autotrend"), nil
}

func main() {
	for i := 0; i < 100000; i++ {
		mgo, _ := ConnectToDB()
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Microsecond)
		mgo.Collection("user").InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
		fmt.Println(i)
	}
}
