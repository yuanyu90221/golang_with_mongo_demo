package main

import (
	"context"
	"fmt"
	"log"
	"time"

	config "github.com/yuanyu90221/golang_with_mongo_demo/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	connectString := config.GetConnectString()
	clientOptions := options.Client().ApplyURI(connectString)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	testDatabase := client.Database("test")
	userCollection := testDatabase.Collection("users")
	userResult, err := userCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "name", Value: "Json Liang"},
			{Key: "age", Value: 18},
			{Key: "skills", Value: bson.A{"programing", "swim", "chess"}},
		},
		bson.D{
			{Key: "name", Value: "nick Liang"},
			{Key: "age", Value: 28},
			{Key: "skills", Value: bson.A{"programing", "swim", "chess", "mountain climbing"}},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("InsertResult", userResult.InsertedIDs)
}
