package domains

import (
	"html/template"
	"os"
)

func GoogleDomainList() {
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"isEven": isEven,
	}).Parse(templ))

	// list of Google domains in domainsGoogle file
	m := GoogleDomainsList

	if err := t.Execute(os.Stdout, m); err != nil {
		panic(err)
	}
}

func BingDomainList() {
	t := template.Must(template.New("").Funcs(template.FuncMap{
		"isEven": isEven,
	}).Parse(templ))

	// list of Bing domains in domainsBing file
	m := BingDomainsList

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