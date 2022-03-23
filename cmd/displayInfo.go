package cmd

import (
	"fmt"
	"html/template"
	"os"
	"strconv"
	"strings"

	build "github.com/Cosiamo/SeaUrchin/build"
	domains "github.com/Cosiamo/SeaUrchin/domains"
)

func DisplayGoogleInfo(searchTerm string, domain string, backoff int) {
	text := "Backoff time was " + strconv.Itoa(backoff) + " seconds"
	fmt.Println(text)

	url, err := build.BuildGoogleUrls(searchTerm, domain, "en", 1, 30)
	if err != nil {
		return
	}
	link := strings.Join(url, "")
	fmt.Println(link)
}

func DisplayBingInfo(searchTerm string, domain string, backoff int) {
	text := "Backoff time was " + strconv.Itoa(backoff) + " seconds"
	fmt.Println(text)

	url, err := build.BuildBingUrls(searchTerm, domain, 1, 30)
	if err != nil {
		return
	}
	link := strings.Join(url, "")
	fmt.Println(link)
}

func GoogleDomainList() {
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"isEven": isEven,
	}).Parse(templ))

	// list of Google domains in domainsGoogle file
	m := domains.GoogleDomainsList

	if err := t.Execute(os.Stdout, m); err != nil {
		panic(err)
	}
}

func BingDomainList() {
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"isEven": isEven,
	}).Parse(templ))

	// list of Bing domains in domainsBing file
	m := domains.BingDomainsList

	if err := t.Execute(os.Stdout, m); err != nil {
		panic(err)
	}
}

func isEven() func() bool {
	e := false
	return func() bool {
		e = !e
		return e
	}
}

// template for displaying the list of domains
const templ = `{{$e := isEven}}
{{- range $k, $v := . -}}
    {{$k}}: {{$v}}
{{end}}`