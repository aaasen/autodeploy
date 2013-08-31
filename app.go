package main

import (
	"github.com/aaasen/dingo"
)

func main() {
	server := dingo.New(dingoConfig)
	server.Run()
}
