package main

import (
	"fmt"
	"ginApi/api"
	"ginApi/storage"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("filed to load .env file: " + err.Error())
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	fmt.Println(connStr)
	// Open a connection to the database
	psqlConnection, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer psqlConnection.Close()

	// Ping the database to verify the connection
	err = psqlConnection.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	log.Println("Connected to the PostgreSQL database!")

	strg := storage.NewStoragePg(psqlConnection)
	apiServer := api.New(&api.RouterOptions{
		Storage: strg,
	})
	port := os.Getenv("PORT")
	err = apiServer.Run(":" + port)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
	log.Print("Server stopped")
}
