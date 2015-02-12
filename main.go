package main

import (
	"github.com/wjsantos/gopr/services"
	"log"
	"os/exec"
)

func main() {

	url := services.NewBitBucket().PRUrl()

	errOpen := exec.Command("open", url).Run()
	if errOpen != nil {
		log.Fatal(errOpen)
	}
}
