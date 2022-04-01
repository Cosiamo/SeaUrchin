package main

import (
	"flag"
	"os"

	cmd "github.com/Cosiamo/SeaUrchin/cmd"
)

func main() {
	// Google
	GoogleCmd := flag.NewFlagSet("g", flag.ExitOnError)
	// Bing
	BingCmd := flag.NewFlagSet("b", flag.ExitOnError)
	// lists supported domains
	DomainsCmd := flag.NewFlagSet("domains", flag.ExitOnError)

	// if user does not input subcommand,
	// returns list of supported subcommands
	if len(os.Args) < 2 {
        cmd.DefaultCase()
    }

	switch os.Args[1] {
	case "g":
		cmd.GoogleCase(GoogleCmd)
	case "b":
		cmd.BingCase(BingCmd)
	case "domains":
		cmd.DomainsCase(DomainsCmd)
	default:
		cmd.DefaultCase()
	}
}