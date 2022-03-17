package cmd

import (
	"flag"
	"fmt"
	"os"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

var searchTerm string

// Switch between Google or Bing depending on subcommand
func SwitchAndCase(GoogleCmd *flag.FlagSet, BingCmd *flag.FlagSet) {
	switch os.Args[1] {
		// Google case
		case "g":
			fmt.Print("Search on Google: ")
			searchTerm = logic.Input()

			GoogleCmd.Parse(os.Args[2:])
			// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
			res, err := logic.GoogleScrape(searchTerm, "com", "en", nil, 1, 30, 10)
			logic.Output(res, err)
		// Bing case
		case "b":
			fmt.Print("Search on Bing: ")
			searchTerm = logic.Input()

			BingCmd.Parse(os.Args[2:])
			// BingScrape(searchTerm, country, proxyString, pages, count, backoff)
			res, err := logic.BingScrape(searchTerm, "com", nil, 1, 30, 10)
			logic.Output(res, err)
		// default case
		default:
			fmt.Println("Expected 'g' or 'b' subcommand")
			os.Exit(1)
		}
}