package cmd

import (
	"flag"
	"fmt"
	"os"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

// Switch between Google or Bing depending on subcommand
func SwitchAndCase(GoogleCmd *flag.FlagSet, BingCmd *flag.FlagSet) {
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
				// need to swap 'Println' for 'log' eventually
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
				// need to swap 'Println' for 'log' eventually
				fmt.Println(err)
			}

		// default case
		default:
			fmt.Println("Expected 'g' or 'b' subcommand")
			os.Exit(1)
		}
}