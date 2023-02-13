package main

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println("Error loading .env file")
	// }

	// log.Println(".env successfully loaded")

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Println(".env successfully loaded")
}

func main() {
	envVariables := make([]string, 0)
	envVariables = append(envVariables, "ZINCSEARCH_USER_ID")
	envVariables = append(envVariables, "ZINCSEARCH_PASSWORD")
	envVariables = append(envVariables, "ZINCSEARCH_ZINCHOST")
	envVariables = append(envVariables, "FALCON_EMAIL_PORT_SERVER")

	for i := 0; i < len(envVariables); i++ {
		value := viper.Get(envVariables[i])
		log.Println("%s: %s", envVariables[i], value)
	}
}
