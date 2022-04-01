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

func GoogleCase(GoogleCmd *flag.FlagSet) {
	proxy := GoogleCmd.String("proxy", "", "Add a proxy string to your search")
	changeDomain := GoogleCmd.String("url", "com", "Choose which region/Google domain you want to search from ('domains' subcommand lists supported domains)")
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
}

func BingCase(BingCmd *flag.FlagSet) {
	proxy := BingCmd.String("proxy", "", "Add a proxy string to your search")
	changeDomain := BingCmd.String("url", "com", "Choose which region/Bing domain you want to search from ('domains' subcommand lists supported domains)")
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
}

func DomainsCase(DomainsCmd *flag.FlagSet) {
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
}

func DefaultCase() {
	fmt.Println("To search, use the subcommands:")
	// google info
	fmt.Fprint(color.Output, "	'", color.YellowString("g"), "' for ")
	fmt.Fprint(color.Output, color.BlueString("G"), color.RedString("o"), color.YellowString("o"), color.BlueString("g"), color.GreenString("l"), color.RedString("e"))
	fmt.Println("")
	fmt.Println("	      or")
	// bing info
	fmt.Fprint(color.Output, "	'", color.YellowString("b"), "' for ")
	fmt.Fprint(color.Output, color.BlueString("B"), color.YellowString("i"), color.BlueString("ng"))
	fmt.Println("")
	fmt.Fprint(color.Output, "Use the flag '", color.CyanString("-help"), "' after the subcommand for more info")
	fmt.Println("")
	fmt.Println(" ")
	fmt.Println("To view the supported regions, use the subcommand:")
	// domains info
	fmt.Fprint(color.Output, "	'", color.YellowString("domains"), "' ")
	fmt.Fprint(color.Output, "with the '", color.CyanString("-g"), "' or '", color.CyanString("-b"), "' flag")
	fmt.Println("")
	fmt.Println(" ")
	os.Exit(1)
}