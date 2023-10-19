package internal

import "strings"

func makeRepoUrl(url string) string {

	if !strings.Contains(url, "http") {
		url = "https://" + url
	}
	return url
}
