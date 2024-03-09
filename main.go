package main

import (
	"fmt"
	// "github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/config"
	// "github.com/zayarhtet/seap-api/src/server/model/dao"
)

func main() {
	fmt.Printf("%s\n", "HELLO WORLD!")
	config.InitEnv()
	config.InitDataCenter()
	config.InitController()
	config.InitRouter()
}
