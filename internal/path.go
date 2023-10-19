package internal

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	} else if empty, _ := isDirEmpty(dirPath); !empty {
		ans := ""
		for ans != "y" && ans != "n" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("dst directory not empty, want clean ? (y/n): ")
			ans, _ = reader.ReadString('\n')
			ans = strings.TrimSpace(ans)
		}
		if ans == "n" {
			panic("abort")
		}

		os.RemoveAll(dirPath)
		os.MkdirAll(dirPath, os.ModePerm)

	}

	return dirPath, nil
}

func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
