package main

import (
	"log"
	config "project/pkg/config"
	di "project/pkg/di"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading the env file \n", err)
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("Couldnt load config:", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server:", diErr)
	} else {
		server.Start()
	}

}
