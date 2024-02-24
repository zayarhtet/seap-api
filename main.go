package main

import (
	"fmt"
	"github.com/zayarhtet/seap-api/src/server/repository"
	"github.com/zayarhtet/seap-api/src/server/config"
	"github.com/zayarhtet/seap-api/src/server/model/dao"
)

func main() {
	fmt.Printf("%s\n", "HELLO WORLD!")
	config.InitEnv()
	var db repository.DataCenter = repository.SeapDataCenter{}
	db.ConnectDatabase()

	var pw string = "HELLO"
	dao.Encrypt(&pw)
	fmt.Println(pw)

}
