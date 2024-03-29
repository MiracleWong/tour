package main

import (
	"github.com/MiracleWong/tour/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal("cmd.Execute err: %v", err)
	}
}
