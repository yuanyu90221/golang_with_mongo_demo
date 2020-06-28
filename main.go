package main

import (
	"context"
	"fmt"
	"log"

	config "github.com/yuanyu90221/golang_with_mongo_demo/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	connectString := config.GetConnectString()
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected To MognoDB with ", connectString)
	cats := client.Database("test").Collection("cats")
	filter := bson.M{}
	var result Cat
	err = cats.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("found single document Name: ", result.Name, ",Age:", result.Age)
}
