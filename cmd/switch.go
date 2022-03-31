package cmd

import (
	"flag"
	"fmt"
	"os"

	client "github.com/Cosiamo/SeaUrchin/client"
	logic "github.com/Cosiamo/SeaUrchin/logic"
	color "github.com/fatih/color"
)

var searchTerm, domain string
var backoff int

// Switch between Google or Bing depending on subcommand
func SwitchAndCase(GoogleCmd *flag.FlagSet, BingCmd *flag.FlagSet, DomainsCmd *flag.FlagSet) {
	switch os.Args[1] {
		// Google case
		case "g":
			proxy := GoogleCmd.String("proxy", "", "Add a proxy string to your search")
			changeDomain := GoogleCmd.String("url", "com", "Choose which Google domain you want to use ('domains' subcommand lists supported domains)")
			showInfo := GoogleCmd.Bool("info", false, "Displays the backoff time and the URL that the results were searched from")
			help := GoogleCmd.Bool("help", false, "Lists flag options")
			rankCmd := GoogleCmd.Bool("rank", false, "Displays the rank of the results next to the title")
			GoogleCmd.Parse(os.Args[2:])

			if *help {
				GoogleCmd.PrintDefaults()
				return
			}

			fmt.Fprint(color.Output, "Search on ", color.BlueString("G"), color.RedString("o"), color.YellowString("o"), color.BlueString("g"), color.GreenString("l"), color.RedString("e"), ": ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			proxyString := client.ConnectProxy(proxy)
			domain = *changeDomain
			// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
			res, err := logic.GoogleScrape(searchTerm, domain, "en", proxyString, 1, 20, backoff)

			logic.Output(res, err, *rankCmd)

			if *showInfo {
				fmt.Fprintln(color.Output, color.GreenString("-----------info-----------"))
				DisplayGoogleInfo(searchTerm, domain, backoff)
			}
		// Bing case
		case "b":
			proxy := BingCmd.String("proxy", "", "Add a proxy string to your search")
			changeDomain := BingCmd.String("url", "com", "Choose which Bing domain you want to use ('domains' subcommand lists supported domains)")
			showInfo := BingCmd.Bool("info", false, "Displays the backoff time and the URL that the results were searched from")
			help := BingCmd.Bool("help", false, "Lists flag options")
			rankCmd := BingCmd.Bool("rank", false, "Displays the rank of the results next to the title")
			BingCmd.Parse(os.Args[2:])

			if *help {
				BingCmd.PrintDefaults()
				return
			}

			fmt.Fprint(color.Output, "Search on ", color.BlueString("B"), color.YellowString("i"), color.BlueString("ng"), ": ")
			searchTerm = logic.Input()
			backoff = logic.Backoff()

			proxyString := client.ConnectProxy(proxy)
			domain = *changeDomain
			// BingScrape(searchTerm, country, proxyString, pages, count, backoff)
			res, err := logic.BingScrape(searchTerm, domain, proxyString, 1, 20, backoff)

			logic.Output(res, err, *rankCmd)

			if *showInfo {
				fmt.Fprintln(color.Output, color.GreenString("-----------info-----------"))
				DisplayBingInfo(searchTerm, domain, backoff)
			}
		// domains case
		case "domains":
			showGoogleDomains := DomainsCmd.Bool("g", false, "Displays all available Google domains")
			showBingDomains := DomainsCmd.Bool("b", false, "Displays all available Bing domains")
			DomainsCmd.Parse(os.Args[2:])

			if !*showGoogleDomains && !*showBingDomains {
				DomainsCmd.PrintDefaults()
				return
			}
			if *showGoogleDomains {
				GoogleDomainList()
				return
			}
			if *showBingDomains {
				BingDomainList()
				return
			}
		// if user inputs invalid subcommand
		default:
			fmt.Println("To search, use the subcommands:")
			fmt.Println("	'g' for Google")
			fmt.Println("	'b' for Bing")
			fmt.Println("To view the supported regions use the subcommand:")
			fmt.Println("	'domains' with the '-g' or '-b' flag")
			os.Exit(1)
		}
}