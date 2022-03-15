package build

import (
	"fmt"
	"strings"

	domains "github.com/Cosiamo/SeaUrchin/domains"
)

// builds the URL
func BuildBingUrls(searchTerm, country string, pages, count int)([]string, error) {
	// toScrape is what is being returned from this function
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if countryCode, found := domains.BingDomains[country]; found {
		for i := 0; i < pages; i++ {
			first := firstParameter(i, count);
			scrapeURL  := fmt.Sprintf("https://bing.com/search?q=%s&first=%d&count=%d%s", searchTerm, first, count, countryCode)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("Country(%s)is currently not supported", country)
		return nil, err
	}
	return toScrape, nil
}

// number is 'i' from the for loop in buildBindUrls
func firstParameter(number, count int) int {
	// need to add 1 because 'i' starts at 0 which Bing does not understand
	if number == 0 {
		return number +1
	}
	return number*count +1
}