package internal

import "strings"

func repoHttpUrl(url string) string {

	if !strings.Contains(url, "http") {
		url = "https://" + url
	}
	return url
}
