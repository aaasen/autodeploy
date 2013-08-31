package main

/*
 * A very limited wrapper for Git
 */

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

type Repo struct {
	path string
}

func NewRepo(path string) (*Repo, error) {
	err := os.Chdir(path)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &Repo{path}, nil
}

func (repo *Repo) Status() {

}

func (repo *Repo) Push(remote, branch string) {

}

func (repo *Repo) Pull(remote, branch string) error {
	cmd := exec.Command("git", "pull", remote, branch)

	out, err := cmd.Output()

	if err != nil {
		return err
	}

	log.Print(string(out))

	return nil
}
