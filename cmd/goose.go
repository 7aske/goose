package main

import (
	"../internal/app/goose"
	"fmt"
)

func main() {
	fmt.Println("starting goose")
	_ = goose.New()
	fmt.Println("exiting goose")
}
