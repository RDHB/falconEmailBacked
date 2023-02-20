package main

import (
	index "falconEmailBackend/scripts"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	log.Println(".env successfully loaded")
}

func main() {
	index.Indexer()
}
