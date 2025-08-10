package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

var mongoClient *mongo.Client

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	if MONGODB_URI == "" {
		log.Fatal("Error: MONGODB_URI not found in .env file")
	}

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient = client
	DB = client.Database("go_crud_db")

	fmt.Println("succesfully connected to MONGODB")

	// unique index creation
	context := context.Background()
	collection := DB.Collection("user")
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(context, indexModel)
	if err != nil {
		log.Fatalf("couldn't create unique index for id field: %v", err)
	}
	fmt.Println("unique index created for id or already exists")

}

// Disconnect function, disconnects the MongoDB connection when the application is closed.
func Disconnect() {
	if mongoClient != nil {
		err := mongoClient.Disconnect(context.Background())
		if err != nil {
			log.Printf("MongoDB bağlantısı kesilirken hata oluştu: %v", err)
		} else {
			fmt.Println("MongoDB bağlantısı güvenli bir şekilde kesildi.")
		}
	}
}