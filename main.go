package main

import (
	"flag"

	cmd "github.com/Cosiamo/SeaUrchin/cmd"
)

func main() {
	// Google
	GoogleCmd := flag.NewFlagSet("g", flag.ExitOnError)
	// Bing
	BingCmd := flag.NewFlagSet("b", flag.ExitOnError)
	// settings
	SettingsCmd := flag.NewFlagSet("settings", flag.ExitOnError)
	// lists supported domains
	DomainsCmd := flag.NewFlagSet("domains", flag.ExitOnError)

	// switches between Google or Bing depending on which subcommand the user inputs
	cmd.SwitchAndCase(GoogleCmd, BingCmd, SettingsCmd, DomainsCmd)
}