package main

import (
	"github.com/douglasdennys/go-mongodb/src/infra/db/mongodb/helpers"
	"github.com/douglasdennys/go-mongodb/src/main/config"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	err := helpers.Connect(os.Getenv("MONGO_URL"), os.Getenv("MONGO_DATABASE"))
	if err != nil {
		panic(err)
	}

	config.StartApp()
}
