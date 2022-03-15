package parse

import (
	"net/http"
	"strings"

	resultModels "github.com/Cosiamo/SeaUrchin/resultModels"
	// external package
	goquery "github.com/PuerkitoBio/goquery"
)

// receive the results from scrapeClientRequest - unstructured data
// creates the struct for the search result with the SearchResult parameters
func BingResultParser(response *http.Response, rank int)([]resultModels.SearchResult, error) {
	// takes response and creates a document from that response
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []resultModels.SearchResult{}
	sel := doc.Find("li.b_algo")
	rank++

	// ranging over all of the node in sel
	for i := range sel.Nodes{
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2")
		descTag := item.Find("div.b_caption p")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := resultModels.SearchResult{
				// ResultRank
				rank,
				// ResultURL
				link,
				// ResultTitle
				title,
				// ResultDesc
				desc,
			}
			results = append(results, result)
			rank ++
		}
	}
	return results, err
}