package services

import "fmt"

type BitBucket struct {
	endpoint string
	git      *Git
}

func NewBitBucket(git *Git) *BitBucket {
	repo := new(BitBucket)
	repo.git = git
	repo.endpoint = "https://bitbucket.org"

	return repo
}

func (repo *BitBucket) PRUrl() (string, error) {
	url := repo.endpoint
	url += repo.buildPath()

	return url, nil
}

func (repo *BitBucket) buildPath() string {
	path := "/" + repo.git.RemoteRepoName()
	path += "/pull-request/new?"
	path += fmt.Sprintf("source=%s", repo.git.RemoteRepoName()+"::"+repo.git.CurrentBranch())

	return path
}
