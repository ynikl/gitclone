package internal

import (
	"log"
	"os"
	"path/filepath"
)

// makeDstPath completes the directory with homedir or pre-set
// directory
func makeDstPath(url string) (string, error) {
	localRoot := localRepoRoot()

	dirPath := filepath.Join(localRoot, url)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
		log.Println("mkdir ", dirPath)
	}
	return dirPath, nil
}
