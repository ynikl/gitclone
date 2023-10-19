package main

import (
	"gitclone/cmd"
	"log"
	"os"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	cloneUrl := os.Args[1]
	if err := cmd.Execute(cloneUrl); err != nil {
		os.Exit(1)
	}
}
