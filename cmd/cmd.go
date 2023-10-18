package cmd

import "gitclone/internal"

func Execute(url string) error {
	return internal.Do(url)
}
