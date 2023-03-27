package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Run() error {
	fmt.Println("starting up application")
	return nil
}

func main() {
	fmt.Println("Go REST API Closure")
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}
}
