package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() (*mongo.Client, context.Context, context.CancelFunc) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	host := os.Getenv("MONGO_HOST")
	log.Println("Connecting on host", host)
	dbUri := fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&ssl=true",
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		host,
	)
	clientOptions := options.Client().
		ApplyURI(dbUri).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		fmt.Println("ping failed!")
		panic(err)
	}
	fmt.Println("ping successful!")
	return client, ctx, cancel
}
