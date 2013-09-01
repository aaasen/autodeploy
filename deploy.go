package autodeploy

import (
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

	repoPath := strings.Join([]string{deployConfig.Dir, userName, repoName}, "/")
	repo, err := NewRepo(repoPath)

	if err != nil {
		return err
	}

	pullErr := repo.Pull("origin", "master")

	if pullErr != nil {
		return pullErr
	}

	return nil
}
