package main

import (
	"github.com/wjsantos/gopr/services"
	"log"
	"os/exec"
)

func main() {

	url, err := new(services.Git).OpenPRUrl()
	if err != nil {
		panic(err)
	}

	if err = exec.Command("open", url).Run(); err != nil {
		log.Fatal(err)
	}
}
