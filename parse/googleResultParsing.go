package parse

import (
	"net/http"
	"strings"

	resultModels "github.com/Cosiamo/SeaUrchin/resultModels"
	// external package
	goquery "github.com/PuerkitoBio/goquery"
)

func GoogleResultParsing(response *http.Response, rank int) ([]resultModels.SearchResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		return nil, err
	}

	results := []resultModels.SearchResult{}
	sel := doc.Find("div.g")
	rank++
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
			rank++
		}
	}

	return results, err
}