package main

import (
	"gitclone/cmd"
	"os"
)

func main() {
	cloneUrl := os.Args[1]
	if err := cmd.Execute(cloneUrl); err != nil {
		os.Exit(1)
	}
}
