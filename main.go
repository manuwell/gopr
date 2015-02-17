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

	errOpen := exec.Command("open", url).Run()
	if errOpen != nil {
		log.Fatal(errOpen)
	}
}
