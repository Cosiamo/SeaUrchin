package cmd

import (
	"flag"
)

var BingInput string

func BingCommand(refinedInput string, BingCmd *flag.FlagSet) {
	b_search := BingCmd.String("", refinedInput, "Search on Bing")
	// converts pointer (*string) to normal value (string)
	BingInput = *b_search
}