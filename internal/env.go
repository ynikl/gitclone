package internal

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

const (
	codeEnv = "code"
	goPath  = "GOPATH"
)

func localRepoRoot() string {
	if codePath := os.Getenv(codeEnv); len(codePath) > 0 {
		return codePath
	}
	if goPath := os.Getenv(goPath); len(goPath) > 0 {
		return goPath
	}
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return home
}
