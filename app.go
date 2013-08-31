package main

import (
	"github.com/aaasen/dingo"
)

const DEPLOY_DIR = "/home/aasen/dev/go/bin/"

func main() {
	server := dingo.New(config)
	server.Run()
}
