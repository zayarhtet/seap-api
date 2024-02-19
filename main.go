package main

import (
	"fmt"
	"github.com/zayarhtet/seap-api/src/server/repository"
)

func main() {
	fmt.Printf("%s\n", "HELLO WORLD!")
	repository.ConnectDatabase()
}
