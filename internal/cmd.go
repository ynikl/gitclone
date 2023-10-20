package internal

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func doCmd(url, dstPath string) error {

	cmd := exec.Command("git", "clone", url, dstPath)
	log.Println(cmd.String())

	done := make(chan struct{})
	go func() {
		defer func() { close(done) }()

		output, err := cmd.CombinedOutput()
		log.Println(string(output))
		if err != nil {
			panic(err)
		}
	}()

	for {
		select {
		default:
			showDoingAnimation()
		case _ = <-done:
			log.Println("done, enjoy!")
			return nil
		}
	}
}

func showDoingAnimation() {
	for _, r := range `-\|/` {
		fmt.Printf("\r%c", r)
		time.Sleep(100 * time.Millisecond)
	}
}
