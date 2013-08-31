package main

import (
	"fmt"
	"strings"
)

func Deploy(userName, repoName string) {
	/*
	 * check name of repo
	 * make a new Repo there
	 * pull down changes
	 * connect to screen
	 * kill process
	 * start process
	 */

	repoPath := strings.Join([]string{DEPLOY_DIR, userName, repoName}, "/")
	repo := NewRepo(repoPath)

	fmt.Println(repo)
}
