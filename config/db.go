package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	mongoString := os.Getenv("MONGOSTRING")
	if mongoString == "" {
		log.Fatal("MONGOSTRING environment variable is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	DB = client
	log.Println("Connected to MongoDB")
}
