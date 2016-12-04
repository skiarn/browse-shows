package templator

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"
)

//New created a new template endpoint.
func New(url BuildAPIURL, render RenderTemplate, validPath string, tmplPaths ...string) *Templator {
	regexp := regexp.MustCompile(validPath)
	tmpls := template.Must(template.ParseFiles(tmplPaths...))
	return &Templator{ValidURLPath: regexp, Templates: tmpls, Render: render, APIURL: url}
}

//BuildAPIURL is a function taking a list parameters previously extracted from
//a request URL and building a specified api url.
type BuildAPIURL func([]string) string

//RenderTemplate takes a io.Reader containing template data and writes it to specified template.
type RenderTemplate func(*template.Template, http.ResponseWriter, string, io.Reader) error

//Templator is handling main flow of http request template rendering.
type Templator struct {
	Render       RenderTemplate
	APIURL       BuildAPIURL
	ValidURLPath *regexp.Regexp
	Templates    *template.Template
}

//IsValid validates a request.
func (p *Templator) IsValid(r *http.Request) (bool, []string) {
	m := p.ValidURLPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		return false, nil
	}
	return true, m
}

//ViewHandler fetches and Renders a view template.
func ViewHandler(w http.ResponseWriter, r *http.Request, params []string, p *Templator) {
	resp, err := http.Get(p.APIURL(params))
	if err != nil {
		fmt.Printf("Unable to call api using request %s \n", p.APIURL(params))
		http.NotFound(w, r)
		return
	}
	err = p.Render(p.Templates, w, "view", resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Request is composed of, http request, http response, scraped url params, templator.
type Request func(http.ResponseWriter, *http.Request, []string, *Templator)

//Handler takes a Request and a regexp used to strip url params from request url.
func Handler(fn Request, p *Templator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isValid, params := p.IsValid(r)
		if !isValid {
			fmt.Printf("Invalid request path %s cannot load content for templator %v \n", r.URL.Path, p)
			http.NotFound(w, r)
			return
		}
		fn(w, r, params, p)
	}
}
