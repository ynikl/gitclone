package internal

import "log"

func Do(url string) error {

	dstPath, err := makeDstPath(url)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("dstPath: ", dstPath)

	srcUrl := makeRepoUrl(url)
	log.Println("srcUrl: ", srcUrl)

	return doCmd(srcUrl, dstPath)
}
