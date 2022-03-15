package cmd

import (
	"flag"
)

var GoogleInput string

func GoogleCommand(refinedInput string, GoogleCmd *flag.FlagSet) {
	g_search := GoogleCmd.String("", refinedInput, "Search on Google")
	// converts pointer (*string) to normal value (string)
	GoogleInput = *g_search
}
