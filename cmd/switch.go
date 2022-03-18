package cmd

import (
	"flag"
	"fmt"
	"os"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

var searchTerm string
var backoff int

// Switch between Google or Bing depending on subcommand
func SwitchAndCase(GoogleCmd *flag.FlagSet, BingCmd *flag.FlagSet, SettingsCmd *flag.FlagSet) {
	switch os.Args[1] {
		// Google case
		case "g":
			fmt.Print("Search on Google: ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			showBackoff := GoogleCmd.Bool("backoff", false, "Displays the backoff time")
			showUrl := GoogleCmd.Bool("url", false, "Displays the URL that the results are from")
			GoogleCmd.Parse(os.Args[2:])

			// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
			res, err := logic.GoogleScrape(searchTerm, "com", "en", nil, 1, 30, backoff)
			logic.Output(res, err)

			if *showBackoff {
				DisplayBackoff(backoff)
			}
			if *showUrl {
				DisplayGoogleUrl(searchTerm)
			}
		// Bing case
		case "b":
			fmt.Print("Search on Bing: ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			showBackoff := BingCmd.Bool("backoff", false, "Displays the backoff time")
			showUrl := BingCmd.Bool("url", false, "Displays the URL that the results are from")
			BingCmd.Parse(os.Args[2:])

			// BingScrape(searchTerm, country, proxyString, pages, count, backoff)
			res, err := logic.BingScrape(searchTerm, "com", nil, 1, 30, backoff)
			logic.Output(res, err)

			if *showBackoff {
				DisplayBackoff(backoff)
			}
			if *showUrl {
				DisplayBingUrl(searchTerm)
			}
		case "settings":
			// googleDomainSet := SettingsCmd.String("gdomain", "com", "Change the Google domain to your region")
			// //bindDomainSet := SettingsCmd.String("bdomain", "com", "Change the Bing domain to your region")
			// SettingsCmd.Parse(os.Args[2:])

			// if len(gdomain) == 3 {

			// } 

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