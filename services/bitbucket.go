package services

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

type BitBucket struct {
	endpoint string
}

func NewBitBucket() *BitBucket {
	repo := new(BitBucket)
	repo.endpoint = "https://bitbucket.org"

	return repo
}

func (repo *BitBucket) PRUrl() string {
	url := repo.endpoint
	url += repo.buildPath()

	return url
}

func (repo *BitBucket) remoteRepoName() string {
	cmd := exec.Command("git", "remote", "-v")

	var remotes bytes.Buffer
	cmd.Stdout = &remotes

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`origin.*:(.+?)\.git \(push\)`)

	return re.FindStringSubmatch(remotes.String())[1]
}

func (repo *BitBucket) currentBranch() string {
	cmd := exec.Command("git", "branch")

	var branches bytes.Buffer
	cmd.Stdout = &branches

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\* (.+)`)

	return re.FindStringSubmatch(branches.String())[1]
}

func (repo *BitBucket) buildPath() string {
	path := "/" + repo.remoteRepoName()
	path += "/pull-request/new?"
	path += fmt.Sprintf("source=%s", repo.remoteRepoName()+"::"+repo.currentBranch())

	return path
}
