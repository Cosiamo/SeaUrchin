package cmd

import (
	"fmt"
	"strconv"
	"strings"

	build "github.com/Cosiamo/SeaUrchin/build"
)

func DisplayGoogleInfo(searchTerm string, domain string, backoff int) {
	text := "Backoff time was " + strconv.Itoa(backoff) + " seconds"
	fmt.Println(text)

	url, err := build.BuildGoogleUrls(searchTerm, domain, "en", 1, 30)
	if err != nil {
		return
	}
	link := strings.Join(url, "")
	fmt.Println(link)
}

func DisplayBingInfo(searchTerm string, domain string, backoff int) {
	text := "Backoff time was " + strconv.Itoa(backoff) + " seconds"
	fmt.Println(text)

	url, err := build.BuildBingUrls(searchTerm, domain, 1, 30)
	if err != nil {
		return
	}
	link := strings.Join(url, "")
	fmt.Println(link)
}