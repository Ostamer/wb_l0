package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	dbConnString string
	kafkaBroker  string
	kafkaTopic   string
)

func initConfig() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnString = os.Getenv("DB_CONN_STRING")
	kafkaBroker = os.Getenv("KAFKA_BROKER")
	kafkaTopic = "orders"

	if dbConnString == "" || kafkaBroker == "" {
		panic("Missing required environment variables")
	}
}
