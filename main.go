package main

import (
	"fmt"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

func main() {
	var input string

	fmt.Println("What do you want to search?")
	fmt.Scanln(&input)

	// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
	res, err := logic.GoogleScrape(input, "com", "en", nil, 1, 30, 10)
	// if no error, range over res var and print a response
	if err  == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}