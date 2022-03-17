package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	logic "github.com/Cosiamo/SeaUrchin/logic"
)

var SeType string

// Switch between Google or Bing depending on subcommand
func SwitchAndCase(GoogleCmd *flag.FlagSet, BingCmd *flag.FlagSet) {
	switch os.Args[1] {
		// Google case
		case "g":
			SeType := "Google"
			
			// sets user input
			input := make([]string, 0)

			// reading user input as a set of lines
			scanner := bufio.NewScanner(os.Stdin)

			fmt.Print("Search on ", SeType, ": ")

			// advancing NewScanner to generate a token
			scanner.Scan()
			// returns the token as a newly allocated string
			text := scanner.Text()
			// appending text, string, to input, string slice
			input = append(input, text)

			// converts the string slice to a single string and concats them with a space
			refinedInput := strings.Join(input, " ")


			GoogleCmd.Parse(os.Args[2:])
			// GoogleScrape(searchTerm, countryCode, languageCode, proxyString, pages, count, backoff)
			res, err := logic.GoogleScrape(refinedInput, "com", "en", nil, 1, 30, 10)

			// if no error, range over res var and print a response
			if err  == nil {
				for _, res := range res {
					fmt.Println(res)
				}
			} else {
				// need to swap 'Println' for 'log' eventually
				fmt.Println(err)
			}

		// Bing case
		case "b":
			SeType := "Bing"

			// sets user input
			input := make([]string, 0)
		
			// reading user input as a set of lines
			scanner := bufio.NewScanner(os.Stdin)
		
			fmt.Print("Search on ", SeType, ": ")
		
			// advancing NewScanner to generate a token
			scanner.Scan()
			// returns the token as a newly allocated string
			text := scanner.Text()
			// appending text, string, to input, string slice
			input = append(input, text)
		
			// converts the string slice to a single string and concats them with a space
			refinedInput := strings.Join(input, " ")
		

			BingCmd.Parse(os.Args[2:])
			// BingScrape(searchTerm, country, proxyString, pages, count, backoff)
			res, err := logic.BingScrape(refinedInput, "com", nil, 1, 30, 10)
			// if no error, range over res var and print a response
			if err == nil {
				for _, res := range res {
					fmt.Println(res)
				}
			} else {
				// need to swap 'Println' for 'log' eventually
				fmt.Println(err)
			}

		// default case
		default:
			fmt.Println("Expected 'g' or 'b' subcommand")
			os.Exit(1)
		}
}