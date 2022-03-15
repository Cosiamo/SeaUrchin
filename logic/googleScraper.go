package logic

import (
	"fmt"

	"net/http"
	"strings"
	"time"

	// internal packages
	client "github.com/Cosiamo/SeaUrchin/client"
	domains "github.com/Cosiamo/SeaUrchin/domains"
	resultModels "github.com/Cosiamo/SeaUrchin/resultModels"

	// external packages
	"github.com/PuerkitoBio/goquery"
)

// creating search queries for Google
func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int)([]string, error) {
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

func googleResultParsing(response *http.Response, rank int)([]resultModels.SearchResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		return nil, err
	}

	results := []resultModels.SearchResult{}
	sel := doc.Find("div.g")
	rank ++
	// iterate over sel and it will have different nodes
	// the nodes have the individual results 
	for i := range sel.Nodes {
		// get the item
		item := sel.Eq(i)
		// find the link tag
		linkTag := item.Find("a")
		// the link is in the link tag
		link, _ := linkTag.Attr("href")
		// find the title tag
		titleTag := item.Find("h3.r")
		// find the description tag
		descTag := item.Find("span.st")
		// get the description from the description tag - convert to text
		desc := descTag.Text()
		// get the title from the title tag - convert to text
		title := titleTag.Text()
		link = strings.Trim(link, " ")

		// if the link is not empty, not null, and doesn't have a prefix
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := resultModels.SearchResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank ++
		}
	}

	return results, err
}

// where the text is posted
func GoogleScrape(searchTerm, countryCode, languageCode string, proxyString interface{}, pages, count, backoff int)([]resultModels.SearchResult, error) {
	results := []resultModels.SearchResult{}
	// count the number of results found for a particular query
	resultCounter := 0
	// creating search queries for Google
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
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
		data, err := googleResultParsing(res, resultCounter)
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