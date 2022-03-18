package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

var searchTerm string
var backoff int

// Switch between Google or Bing depending on subcommand
func SwitchAndCase(GoogleCmd *flag.FlagSet, BingCmd *flag.FlagSet) {
	switch os.Args[1] {
		// Google case
		case "g":
			fmt.Print("Search on Google: ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			GoogleCmd.Parse(os.Args[2:])
			// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
			res, err := logic.GoogleScrape(searchTerm, "com", "en", nil, 1, 30, backoff)
			logic.Output(res, err)
			test := "Backoff time was " + strconv.Itoa(backoff) + " seconds"
			fmt.Println(test)
		// Bing case
		case "b":
			fmt.Print("Search on Bing: ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			BingCmd.Parse(os.Args[2:])
			// BingScrape(searchTerm, country, proxyString, pages, count, backoff)
			res, err := logic.BingScrape(searchTerm, "com", nil, 1, 30, backoff)
			logic.Output(res, err)
			test := "Backoff time was " + strconv.Itoa(backoff) + " seconds"
			fmt.Println(test)
		// if user inputs invalid subcommand
		default:
			fmt.Println("To search, use the subcommands:")
			fmt.Println("    'g' for Google")
			fmt.Println("    'b' for Bing")
			// use this for when settings are implemted for their various flags
			flag.PrintDefaults()
			os.Exit(1)
		}
}