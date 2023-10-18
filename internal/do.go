package internal

import "log"

func Do(url string) error {

	dstPath, err := makeDstPath(url)
	if err != nil {
		log.Fatalln(err)
	}
	srcUrl := repoHttpUrl(url)
	return doCmd(srcUrl, dstPath)
}
