package main

import (
	"github.com/EraldCaka/prizz-backend/router"
	"log"
)

func main() {
	router.NewRouter()
	if err := router.Start(":5555"); err != nil {
		log.Fatal(err)
	}

}
