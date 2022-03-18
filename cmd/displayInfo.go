package cmd

import (
	"fmt"
	"strconv"
	"strings"

	build "github.com/Cosiamo/SeaUrchin/build"
)

func DisplayBackoff(backoff int) {
	text := "Backoff time was " + strconv.Itoa(backoff) + " seconds"
	fmt.Println(text)
}

func DisplayGoogleUrl(searchTerm string) {
	url, err := build.BuildGoogleUrls(searchTerm, "com", "en", 1, 30)
	if err != nil {
		return
	}
	text := strings.Join(url, "")
	fmt.Println(text)
}

func DisplayBingUrl(searchTerm string) {
	url, err := build.BuildBingUrls(searchTerm, "com", 1, 30)
	if err != nil {
		return
	}
	text := strings.Join(url, "")
	fmt.Println(text)
}