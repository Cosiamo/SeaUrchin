package main

import (
	"fmt"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

func main() {
	// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
	res, err := logic.GoogleScrape("github", "com", "en", nil, 1, 30, 10)
	// if no error, range over res var and print a response
	if err  == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}