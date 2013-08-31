package main

import (
	"fmt"
	"net/http"
)

type RepoUpdateController struct{}

func (c RepoUpdateController) Respond(w http.ResponseWriter, r *http.Request, data map[string]string) {
	fmt.Fprint(w, "hey")
}
