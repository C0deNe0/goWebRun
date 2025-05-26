package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("URI is empty")
	}

	//getting a client
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	//disconnecting from the client
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//selecting the collection
	coll := client.Database("mydb").Collection("cars")

	var title = "Back to future"
	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{
		{Key: "title", Value: title},
	}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s", title)
		return
	}
	if err!=nil{
		
	}

}
