package main

import (
	"github.com/aaasen/dingo"
	"os"
)

var dingoConfig = dingo.Config{
	Port:        "8080",
	TemplateDir: "/",
	StaticDir:   "/",
	Routes:      routes,
}

type DeployConfig struct {
	Dir string
}

var deployConfig = &DeployConfig{os.ExpandEnv("$HOME/dev/go/bin")}
