package internal

import (
	"os"
	"os/exec"
	"strings"

	"github.com/mitchellh/go-homedir"
)

const (
	codeEnv = "CODE"
	goPath  = "GOPATH"
)

func localRepoRoot() string {

	if codePath := os.Getenv(codeEnv); len(codePath) > 0 {
		return codePath
	}
	if goPath := os.Getenv(goPath); len(goPath) > 0 {
		return goPath
	}
	if goEnvGoPath := goEnvGOPATH(); len(goEnvGoPath) > 0 {
		return goEnvGoPath
	}

	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return home
}

func goEnvGOPATH() string {

	cmd := exec.Command("go", "env", "GOPATH")
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(output))
}
