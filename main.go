package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

func main() {
    input := make([]string, 0)

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("What do you want to search? : ")

    scanner.Scan()
    text := scanner.Text()
	input = append(input, text)

	refinedInput := strings.Join(input, "+")

	// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
	res, err := logic.GoogleScrape(refinedInput, "com", "en", nil, 1, 30, 10)
	// if no error, range over res var and print a response
	if err  == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}