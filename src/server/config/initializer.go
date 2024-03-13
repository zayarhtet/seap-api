package config

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/zayarhtet/seap-api/src/server/controller"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/router"
)

func InitEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println("Environment variables have been initialised.")
}

func InitDataCenter() {
	repository.Init()
	log.Println("Data Center has been initialised.")
}

func InitRouter() {
	router.Init()
	log.Println("Routers has been initialised.")
}

func InitController() {
	controller.Init()
	log.Println("Constrollers has been initialised.")
}
