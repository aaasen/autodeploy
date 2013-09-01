package autodeploy

import (
	"github.com/aaasen/dingo"
	"os"
)

type DeployConfig struct {
	Dir string
}

var dingoConfig = dingo.Config{
	Port:        "8080",
	TemplateDir: "/",
	StaticDir:   "/",
	Routes:      routes,
}

var deployConfig = &DeployConfig{os.ExpandEnv("$HOME/dev/go/bin")}
