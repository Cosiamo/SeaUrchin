package logic

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	resultModels "github.com/Cosiamo/SeaUrchin/resultModels"
)

func Input() string {
	// sets user input
	input := make([]string, 0)

	// reading user input as a set of lines
	scanner := bufio.NewScanner(os.Stdin)

	// advancing NewScanner to generate a token
	scanner.Scan()
	// returns the token as a newly allocated string
	text := scanner.Text()
	// appending text, string, to input, string slice
	input = append(input, text)

	// converts the string slice to a single string and concats them with a space
	refinedInput := strings.Join(input, " ")
	return refinedInput
}

func Output(res []resultModels.SearchResult, err error) {
	// if no error, range over res var and print a response
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	} else {
		// need to swap 'Println' for 'log' eventually
		fmt.Println(err)
	}
}