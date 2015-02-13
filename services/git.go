package services

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"
)

type Git struct{}

func (git *Git) RemoteRepoName() string {
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
