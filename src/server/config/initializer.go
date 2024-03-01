package config

import (
	"github.com/joho/godotenv"
	"github.com/zayarhtet/seap-api/src/server/controller"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/router"
	"log"
)

func InitEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func InitDataCenter() {
	repository.Init()
}

func InitRouter() {
	router.Init()
}

func InitController() {
	controller.Init()
}
