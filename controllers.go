package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RepoUpdateController struct{}

func (c RepoUpdateController) Respond(w http.ResponseWriter, r *http.Request, data map[string]string) {
	rawPushInfo := r.FormValue("payload")

	var pushInfo interface{}
	err := json.Unmarshal([]byte(rawPushInfo), &pushInfo)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	repoInfo := pushInfo.(map[string]interface{})["repository"].(map[string]interface{})
	userName := repoInfo["owner"].(map[string]interface{})["name"]
	repoName := repoInfo["name"]

	Deploy(userName.(string), repoName.(string))
}
