package main

import (
	"User-service/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database, err := db.ConnectDB()
	if err != nil {
		log.Println(err)
		log.Fatal("DB connection could not be set")
	}
	storage:=db.NewStorage(database)

}
