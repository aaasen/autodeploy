package main

import (
	"fmt"
	"strings"
)

func Deploy(userName, repoName string) error {
	/*
	 * check name of repo
	 * make a new Repo there
	 * pull down changes
	 * connect to screen
	 * kill process
	 * start process
	 */

	repoPath := strings.Join([]string{DEPLOY_DIR, userName, repoName}, "/")
	repo, err := NewRepo(repoPath)

	if err != nil {
		return err
	}

	repo.Pull("origin", "master")

	fmt.Println(repo)

	return nil
}
