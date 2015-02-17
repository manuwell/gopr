package services

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type Git struct{}

func (git *Git) OpenPRUrl() (string, error) {
	if strings.Contains("github", git.ServiceName()) {
		return NewGithub(git).PRUrl()
	} else if strings.Contains("bitbucket", git.ServiceName()) {
		return NewBitBucket(git).PRUrl()
	} else {
		return "", errors.New("Service not supported. Only GitHub and BitBucket")
	}
}

func (git *Git) RemoteRepoName() string {
	cmd := exec.Command("git", "remote", "-v")

	var remotes bytes.Buffer
	cmd.Stdout = &remotes

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`origin.*[\/:](.+?\/.+?)\.git \(push\)`)

	return re.FindStringSubmatch(remotes.String())[1]
}

func (git *Git) CurrentBranch() string {
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

func (git *Git) ServiceName() string {
	// @todo: Must be refactored with the RemoteRepoName
	cmd := exec.Command("git", "remote", "-v")

	var remotes bytes.Buffer
	cmd.Stdout = &remotes

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`@(\w+)`)

	return re.FindStringSubmatch(remotes.String())[1]
}
