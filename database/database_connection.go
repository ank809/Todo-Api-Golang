package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBinstance()

func DBinstance() *mongo.Client {
	url := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println(err)

	}
	return client
}
func OpenCollection(client *mongo.Client, c string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("mongo-go").Collection(c)
	return collection
}
