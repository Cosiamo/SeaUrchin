package main

import (
	// go packages
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	// internal packages
	logic "github.com/Cosiamo/SeaUrchin/logic"
)

func main() {
	var GoogleInput string

	// sets user input
    input := make([]string, 0)

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("What do you want to search? : ")

    scanner.Scan()
    text := scanner.Text()
	input = append(input, text)

	// converts the string slices to a single string and concats them with "+"
	refinedInput := strings.Join(input, "+")

	// subcommands - flags
	GoogleCmd := flag.NewFlagSet("g", flag.ExitOnError)
	g_search := GoogleCmd.String("", refinedInput, "Search on Google")

	// converts pointer (*string) to normal value (string)
	GoogleInput = *g_search

	switch os.Args[1] {
	// Google case
	case "g":
		GoogleCmd.Parse(os.Args[2:])
		// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
		res, err := logic.GoogleScrape(GoogleInput, "com", "en", nil, 1, 30, 10)
		// if no error, range over res var and print a response
		if err  == nil {
			for _, res := range res {
				fmt.Println(res)
			}
		}
	// default case
	default:
		fmt.Println("Expected g subcommand")
		os.Exit(1)
	}


}