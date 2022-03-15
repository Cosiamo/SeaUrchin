package logic

import (
	"time"

	build "github.com/Cosiamo/SeaUrchin/build"
	client "github.com/Cosiamo/SeaUrchin/client"
	parse "github.com/Cosiamo/SeaUrchin/parse"
	resultModels "github.com/Cosiamo/SeaUrchin/resultModels"
)

// assembles the results from all the other functions
func BingScrape(searchTerm, country string, proxyString interface{}, pages, count, backoff int)([]resultModels.SearchResult, error) {
	results := []resultModels.SearchResult{}

	bingPages, err := build.BuildBingUrls(searchTerm, country, pages, count)

	if err != nil {
		return nil, err
	}

	for _, page := range bingPages{
		rank := len(results)
		res, err := client.ScrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := parse.BingResultParser(res, rank)
		if err != nil {
			return nil, err
		}
		// range over the data and append to results
		// create a slice of result and send it back
		for _, result := range data {
			results = append(results, result)
		}
		// optional but a best practice  when scraping something or making requests
		// ideally want to use backoff to randomize the time duration between when which you make the requests
		time.Sleep(time.Duration(backoff)*time.Second)
	}
	return results, nil
}