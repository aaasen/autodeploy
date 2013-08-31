package main

/*
 * A very limited wrapper for Git
 */

import ()

type Repo struct {
	path string
}

func NewRepo(path string) *Repo {
	return &Repo{path}
}

func (repo *Repo) Push(remote, branch string) {

}

func (repo *Repo) Pull(remote, branch string) {

}
