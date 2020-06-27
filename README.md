# golang_with_mongodb

[golang with docker](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents)

## introduction

this is an exmple for golang to make connect to the mongodb

use the mongo db offical library [mongo](go.mongodb.org/mongo-driver/mongo)

## initial
```script===
go mod init github.com/${your github account name}/${github repo name}
go get -u github.com/joho/godotenv
go get -u go.mongodb.org/mongo-driver/mongo
go get -u go.mongodb.org/mongo-driver/bson
```
## connect setup
```golang==
clientOptions := options.Client().ApplyURI(connectString)
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil {
    log.Fatal(err)
}
```