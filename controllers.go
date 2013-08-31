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
		log.Panic(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	repoInfo := pushInfo.(map[string]interface{})["repository"].(map[string]interface{})
	userName := repoInfo["owner"].(map[string]interface{})["name"]
	repoName := repoInfo["name"]

	deployErr := Deploy(userName.(string), repoName.(string))

	if deployErr != nil {
		log.Panic(deployErr)
		http.Error(w, deployErr.Error(), http.StatusBadRequest)
	}

	log.Printf("completed build of %s/%s\n", userName, repoName)
}
