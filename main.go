package main

import (
	// go packages
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	cmd "github.com/Cosiamo/SeaUrchin/cmd"
)

func main() {
	// sets user input
    input := make([]string, 0)

	// reading user input as a set of lines
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("What do you want to search? : ")

	// advancing NewScanner to generate a token
    scanner.Scan()
	// returns the token as a newly allocated string
    text := scanner.Text()
	// appending text, string, to input, string slice
	input = append(input, text)

	// converts the string slice to a single string and concats them with a space
	refinedInput := strings.Join(input, " ")

	// Google
	GoogleCmd := flag.NewFlagSet("g", flag.ExitOnError)
	cmd.GoogleCommand(refinedInput, GoogleCmd)

	// Bing
	BingCmd := flag.NewFlagSet("b", flag.ExitOnError)
	cmd.BingCommand(refinedInput, BingCmd)

	// switches between Google or Bing depending on which subcommand the user inputs
	cmd.SwitchAndCase(GoogleCmd, BingCmd)
}