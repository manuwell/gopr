package services

import "fmt"

type Github struct {
	endpoint string
	git      *Git
}

func NewGithub(git *Git) *Github {
	repo := new(Github)
	repo.git = git
	repo.endpoint = "https://github.com"

	return repo
}

func (repo *Github) PRUrl() (string, error) {
	url := repo.endpoint
	url += repo.buildPath()

	return url, nil
}

func (repo *Github) buildPath() string {

	path := "/" + repo.git.RemoteRepoName()
	path += fmt.Sprintf("/compare/%s", repo.git.CurrentBranch())

	return path
}
