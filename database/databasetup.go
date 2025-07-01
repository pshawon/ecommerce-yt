package database

import (
	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func DBSet() *mongo.Client{
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeOut(context.Background(), 10 *time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err!= nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err!= nil {
		log.Println("Failed to connect to mongodb: (")
		return nil
	}

	fmt.Println("Successfully connected to mongodb")
	return client
}

func UserData( client *moongo.Client, collectionName string) *mongo.Collection{

}

func ProductData(client *moongo.Client, collectionName string){

}