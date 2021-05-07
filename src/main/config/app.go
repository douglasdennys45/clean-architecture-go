package config

import (
	"fmt"
	"net/http"
)

func StartApp() {
	Routes()
	fmt.Println("started server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}