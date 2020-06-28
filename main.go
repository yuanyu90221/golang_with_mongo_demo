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
	// var result Cat
	var resultList []Cat
	// err = cats.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	cur, errAll := cats.Find(context.Background(), filter)
	if errAll != nil {
		log.Fatal(errAll)
	}
	for cur.Next(context.Background()) {
		var cat Cat
		errTemp := cur.Decode(&cat)
		if errTemp != nil {
			log.Fatal(errTemp)
		}
		resultList = append(resultList, cat)
	}
	if len(resultList) > 0 {
		deleteResult, err1 := cats.DeleteMany(context.TODO(), filter)
		if err1 != nil {
			log.Fatal(err1)
		}
		log.Println(*&deleteResult.DeletedCount)
	}
	err = client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("found single document Name: ", result.Name, ",Age:", result.Age)
	fmt.Println(resultList)
}
