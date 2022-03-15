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

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("What do you want to search? : ")

    scanner.Scan()
    text := scanner.Text()
	input = append(input, text)

	// converts the string slices to a single string and concats them with a space
	refinedInput := strings.Join(input, " ")

	// Google
	GoogleCmd := flag.NewFlagSet("g", flag.ExitOnError)
	cmd.GoogleCommand(refinedInput, GoogleCmd)

	// Bing
	BingCmd := flag.NewFlagSet("b", flag.ExitOnError)
	cmd.BingCommand(refinedInput, BingCmd)

	cmd.SwitchAndCase(GoogleCmd, BingCmd)
}