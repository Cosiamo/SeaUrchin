package client

import (
	"net/http"
	"net/url"
)

func GetScrapeClient(proxyString interface{}) *http.Client {
	switch V := proxyString.(type){
	// if a string is passed in the proxy
	case string:
		// parse the proxyString
		proxyUrl, _ := url.Parse(V)
		// proxyUrl becomes the Proxy value which becomes the Transport value that becomes the client
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	// if a string is NOT passed in the proxy
	default:
		return &http.Client{}
	}
}