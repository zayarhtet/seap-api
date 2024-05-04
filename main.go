package main

import (
	"log"

	"github.com/zayarhtet/seap-api/src/engine"
	"github.com/zayarhtet/seap-api/src/server/config"
)

func main() {
	log.Printf("%s\n", "HELLO WORLD!")
	engine.Init()
	config.InitEnv()
	config.InitDataCenter()
	config.InitController()
	config.InitRouter()

}
