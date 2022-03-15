package logic

import (
	"time"

	build "github.com/Cosiamo/SeaUrchin/build"
	client "github.com/Cosiamo/SeaUrchin/client"
	parse "github.com/Cosiamo/SeaUrchin/parse"
	resultModels "github.com/Cosiamo/SeaUrchin/resultModels"
)

// where the text is posted
func GoogleScrape(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int)([]resultModels.SearchResult, error) {
	results := []resultModels.SearchResult{}
	// count the number of results found for a particular query
	resultCounter := 0
	// creating search queries for Google
	googlePages, err := build.BuildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}

	// scrape, with ScrapeClientRequest, through googlePages query one by one 
	for _, page := range googlePages {
		// ScrapeClientRequest will make the request to the query googlePages
		res, err := client.ScrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		// sending the response that was received after the client was scraped and sending the resultCounter
		data, err := parse.GoogleResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data {
			results = append(results, result)
		}
		// will sleep after every iteration of the for loop above
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}