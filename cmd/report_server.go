package main

import (
	"github.com/EraldCaka/prizz-backend/router"
	"log"
)

func main() {
	router.NewReportRouter()
	if err := router.StartReportServer(":5550"); err != nil {
		log.Fatal(err)
	}

}
