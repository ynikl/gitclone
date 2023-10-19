package internal

import (
	"fmt"
	"log"
	"os/exec"
	"sync/atomic"
	"time"
)

func doCmd(url, dstPath string) error {

	cmd := exec.Command("git", "clone", url, dstPath)
	log.Println(cmd.String())

	value := &atomic.Value{}
	value.Store(int64(0))
	go func() {
		defer func() { value.Store(int64(1)) }()

		output, err := cmd.CombinedOutput()
		log.Println(string(output))
		if err != nil {
			panic(err)
		}
	}()

	showDoingAnimation(value)
	log.Println("done, enjoy!")

	return nil
}

func showDoingAnimation(value *atomic.Value) {
	for value.Load().(int64) == 0 {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
