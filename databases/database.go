package database

import (
 "context"
 "fmt"
 "log"
 "time"
 "go.mongodb.org/mongo-driver/mongo"
 "go.mongodb.org/mongo-driver/mongo/options"
 "os"
 "github.com/joho/godotenv"
)

func ConnectDB() *mongo.Client {
	godotenv.Load(".env")
	password := os.Getenv("App_password")
 Mongo_URL := "mongodb+srv://craftsintayehu:"+password+"@cluster0.4tk0ke9.mongodb.net/?retryWrites=true&w=majority" // Replace this with your connection string.
 client, err := mongo.NewClient(options.Client().ApplyURI(Mongo_URL))

 if err != nil {
 log.Fatal(err)
 }

 ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
 err = client.Connect(ctx)
 defer cancel()

 if err != nil {
 log.Fatal(err)
 }

 fmt.Println("Connected to mongoDB")
 return client
}