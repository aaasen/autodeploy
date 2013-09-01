package autodeploy

import (
	"github.com/aaasen/dingo"
)

func Start() {
	server := dingo.New(dingoConfig)
	server.Run()
}
