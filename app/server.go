package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/spf13/viper"
	"github.com/joho/godotenv"

	falconEmailServer "falconEmailBackend/api/router"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	log.Println(".env successfully loaded")
}

func main() {
	port := fmt.Sprintf("%s", os.Getenv("FALCON_EMAIL_PORT_SERVER"))
	log.Printf("Serving on port%s", port)
	http.ListenAndServe(port, falconEmailServer.InitializeZincSearchRouter()) //revisar y cambiar por lo comentado en la versión final
	//	server := &http.Server{
	//		Addr:              port,
	//		Handler:           falconEmailServer.InitializeZincSearchRouter(),
	//		ReadHeaderTimeout: 3 * time.Second,
	//	}
	//
	// err := server.ListenAndServe()
	//
	//	if err != nil {
	//		log.Fatalf(fmt.Sprintf("Error initializing the server %s", err))
	//	}
}
