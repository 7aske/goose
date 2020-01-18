package main

import (
	"../internal/app/goose"
	"fmt"
	"log"
)

func main() {
	fmt.Println("starting goose")
	err := goose.New()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("exiting goose")
}
