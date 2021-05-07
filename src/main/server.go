package main

import (
	"github.com/douglasdennys/clean-architecture-go/src/drivers/database/mongodb/shared"
	"github.com/douglasdennys/clean-architecture-go/src/main/config"
)

func main() {
	err := shared.Connect("mongodb://127.0.0.1:27017", "apps")
	if err != nil {
		panic(err)
	}

	config.StartApp()
}