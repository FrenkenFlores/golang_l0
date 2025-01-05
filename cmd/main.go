package main

import (
	"log"

	gol0 "github.com/FrenkenFlores/golang_l0"
)

func main() {
	port := "8080"
	if err := gol0.NewServer(port); err != nil {
		log.Fatalf("Error occurred while setting up the server: %s", err.Error())
	}
}
