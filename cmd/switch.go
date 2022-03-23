package cmd

import (
	"flag"
	"fmt"
	"os"

	domains "github.com/Cosiamo/SeaUrchin/domains"
	logic "github.com/Cosiamo/SeaUrchin/logic"
)

var searchTerm string
var backoff int
var domain string

// Switch between Google or Bing depending on subcommand
func SwitchAndCase(GoogleCmd *flag.FlagSet, BingCmd *flag.FlagSet, SettingsCmd *flag.FlagSet, DomainsCmd *flag.FlagSet) {
	switch os.Args[1] {
		// Google case
		case "g":
			fmt.Print("Search on Google: ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			changeDomain := GoogleCmd.String("url", "com", "Choose which Google domain you want to use")
			showInfo := GoogleCmd.Bool("info", false, "Displays the backoff time and the URL that the results were searched from")
			GoogleCmd.Parse(os.Args[2:])

			domain = *changeDomain
			// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
			res, err := logic.GoogleScrape(searchTerm, domain, "en", nil, 1, 30, backoff)
			logic.Output(res, err)

			if *showInfo {
				DisplayGoogleInfo(searchTerm, domain, backoff)
			}
		// Bing case
		case "b":
			fmt.Print("Search on Bing: ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			changeDomain := BingCmd.String("url", "com", "Choose which Bing domain you want to use")
			showInfo := BingCmd.Bool("info", false, "Displays the backoff time and the URL that the results were searched from")
			BingCmd.Parse(os.Args[2:])

			domain = *changeDomain
			// BingScrape(searchTerm, country, proxyString, pages, count, backoff)
			res, err := logic.BingScrape(searchTerm, domain, nil, 1, 30, backoff)
			logic.Output(res, err)

			if *showInfo {
				DisplayBingInfo(searchTerm, domain, backoff)
			}
		case "domains":
			showGoogleDomains := DomainsCmd.Bool("g", false, "Displays all available Google domains")
			showBingDomains := DomainsCmd.Bool("b", false, "Displays all available Bing domains")
			DomainsCmd.Parse(os.Args[2:])

			if !*showGoogleDomains && !*showBingDomains {
				fmt.Println("To view the list of supported domains use the flag '-g' for Google or '-b' for Bing")
				return
			}
			if *showGoogleDomains {
				domains.GoogleDomainList()
				return
			}
			if *showBingDomains {
				domains.BingDomainList()
				return
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
			fmt.Println("'g' for Google")
			fmt.Println("'b' for Bing")
			fmt.Println("Use the '-info' flag to see the backoff time and the link that was used to fetch the results")
			fmt.Println("Use the '-url' flag to change the region you're searching in")
			fmt.Println("To view the supported regions use the subcommand 'domains' with the '-g' or '-b' flag")
			// use this for when settings are implemted for their various flags
			flag.PrintDefaults()
			os.Exit(1)
		}
}