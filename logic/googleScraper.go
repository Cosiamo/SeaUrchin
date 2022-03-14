package logic

import (
	"fmt"

	"math/rand"
	// for http requests
	"net/http"
	"net/url"
	"strings"
	"time"

	domains "github.com/Cosiamo/SeaUrchin/domains"
	userAgents "github.com/Cosiamo/SeaUrchin/browserEngines"
	resultModels "github.com/Cosiamo/SeaUrchin/resultModels"

	// help with scrapping from google
	"github.com/PuerkitoBio/goquery"
)

// this is used so that Google thinks the requests are coming from different browsers
// need this so that Google doesn't think anything shady is going on
func randomUserAgent() string {
	// select a random number
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents.UserAgents)
	// to access a particular value in a slice there needs to be an index passed into an array
	return userAgents.UserAgents[randNum]
}

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

func getScrapeClient(proxyString interface{}) *http.Client {

	switch v := proxyString.(type){

	// if a string is passed in the proxy
	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	// if a string is NOT passed in the proxy
	default:
		return &http.Client{}
	}
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

	// scrape, with scrapeClientRequest, through googlePages query one by one 
	for _, page := range googlePages {
		// scrapeClientRequest will make the request to the query googlePages
		res, err := scrapeClientRequest(page, proxyString)
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

// only receiving one page from for loop
func scrapeClientRequest(searchURL string, proxyString interface{})(*http.Response, error) {
	// if there's something in the proxy string it will create a baseClient 
	// else it will return a default baseClient
	baseClient := getScrapeClient(proxyString)
	// a GET request to the searchURL from buildGoogleUrls function
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", randomUserAgent())

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