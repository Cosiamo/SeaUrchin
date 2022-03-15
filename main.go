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
	var BingInput string

	// sets user input
    input := make([]string, 0)

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("What do you want to search? : ")

    scanner.Scan()
    text := scanner.Text()
	input = append(input, text)

	// converts the string slices to a single string and concats them with "+"
	refinedInput := strings.Join(input, " ")

	// subcommands - flags
	// Google
	GoogleCmd := flag.NewFlagSet("g", flag.ExitOnError)
	g_search := GoogleCmd.String("", refinedInput, "Search on Google")
	// converts pointer (*string) to normal value (string)
	GoogleInput = *g_search

	// Bing
	BingCmd := flag.NewFlagSet("b", flag.ExitOnError)
	b_search := BingCmd.String("", refinedInput, "Search on Bing")
	// converts pointer (*string) to normal value (string)
	BingInput = *b_search

	// Switch between Google or Bing depending on flag
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
		} else {
			// a lot of projects use 'log' for errors
			fmt.Println(err)
		}
	// Bing case
	case "b":
		BingCmd.Parse(os.Args[2:])
		// BingScrape(searchTerm, country, proxyString, pages, count, backoff)
		res, err := logic.BingScrape(BingInput, "com", nil, 1, 30, 10)
		// if no error, range over res var and print a response
		if err == nil {
			for _, res := range res {
				fmt.Println(res)
			}
		} else {
			// a lot of projects use 'log' for errors
			fmt.Println(err)
		}
	// default case
	default:
		fmt.Println("Expected g subcommand")
		os.Exit(1)
	}
}