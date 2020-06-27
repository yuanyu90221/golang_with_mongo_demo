package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	// load
	DB_HOST := os.Getenv("MONGO_HOST")
	DB_PORT := os.Getenv("PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USERBAME := os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	DB_PASSWD := os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	connectString := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", DB_USERBAME, DB_PASSWD, DB_HOST, DB_PORT, DB_NAME)
	log.Println("connectString", connectString)
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
	fmt.Println("found single document: ", result)
}
