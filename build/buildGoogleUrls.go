package build

import (
	"fmt"
	"strings"

	domains "github.com/Cosiamo/SeaUrchin/domains"
)

// creating search queries for Google
func BuildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int)([]string, error) {
	toScrape := []string{}

	// refining the search term
	searchTerm = strings.Trim(searchTerm, " ")
	// + is used in the URL instead of spaces
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)

	// find the realvent Google domain (URL) for the country code
	// helps find the index for the countryCode
	if googleBase, found := domains.GoogleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			// "%s is googlebase, %s is searchTerm, num=%d is count, hl=%s is language code, start=%d is start
			scrapeURL := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
			toScrape = append(toScrape, scrapeURL)
		}
	} else {
		err := fmt.Errorf("country (%s) is currently not supported", countryCode)
		return nil, err
	}
	return toScrape, nil
}