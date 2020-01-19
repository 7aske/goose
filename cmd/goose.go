package main

import (
	"../internal/app/goose"
	"log"
)

func main() {
	log.Println("starting goose")
	goose.New()
	log.Println("exiting goose")
}
