package client

import (
	"fmt"
	"net/http"

	userAgents "github.com/Cosiamo/SeaUrchin/userAgents"
)

// only receiving one page from for loop
func ScrapeClientRequest(searchURL string, proxyString interface{})(*http.Response, error) {
	// if there's something in the proxy string it will create a baseClient 
	// else it will return a default baseClient
	baseClient := GetScrapeClient(proxyString)
	// a GET request to the searchURL from buildGoogleUrls function
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", userAgents.RandomUserAgent())

	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("scraper received a non-200 status code suggesting a ban")
		// if the status code is not 200, returning nil for res and err for err
		return nil, err
	}

	if err != nil {
		// if there's an error, returning nil for res and err for error
		return nil, err
	}
	
	// returning res to http.Response and nil to error
	return res, nil
}