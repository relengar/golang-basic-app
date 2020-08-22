package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB is a database instace of connection to the database
var DB *mongo.Database

// Connect to the database
func Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(getConnectionString()))
	if err != nil {
		log.Fatalf("Mongo client setup error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Mongo connection error: %v", err)
	}
	// defer client.Disconnect(ctx)

	DB = client.Database("recipe_simple")
	fmt.Println("Connected to db")
	recipesCollection = DB.Collection("recipes")
}

func getConnectionString() string {
	user := getEnvVar("MONGO_USER")
	password := getEnvVar("MONGO_PASSWORD")
	address := getEnvVar("MONGO_ADDRESS")
	port := getEnvVar("MONGO_PORT")
	dbName := getEnvVar("MONGO_DB_NAME")
	uriParts := []string{"mongodb://", user, ":", password, "@", address, ":", port, "/", dbName, "?retryWrites=false"}
	return strings.Join(uriParts, "")
}

func getEnvVar(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Missing %v environment variable", key)
	}
	return value
}
