package templator_test

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/skiarn/browse-shows/templator"
)

type Test struct {
	Name        string
	LuckyNumber int64
}

func TestViewHandler(t *testing.T) {
	regexp := regexp.MustCompile("^/(view)/([a-zA-Z0-9_-]+)$")
	tmpl := template.Must(template.New("view.html").Parse(`<h1>{{.Name}}<h2><p>I'm lucky {{.LuckyNumber}}</p>`))

	//setup remote api
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"Name":"Alice","LuckyNumber":41}`)
	}))
	defer apiServer.Close()
	//implements templator.BuildAPIURL
	url := func(params []string) string {
		return apiServer.URL + "/api/test/programs/" + params[2]
	}

	//Render implements templator.RenderTemplate
	render := func(tmpl *template.Template, w http.ResponseWriter, tmplName string, body io.Reader) error {
		var info Test
		if err := json.NewDecoder(body).Decode(&info); err != nil {
			return err
		}
		return tmpl.ExecuteTemplate(w, tmplName+".html", info)
	}
	templatorTest := &templator.Templator{ValidURLPath: regexp, Templates: tmpl, Render: render, APIURL: url}

	req, err := http.NewRequest("GET", "/view/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(templator.Handler(templator.ViewHandler, templatorTest))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `<h1>Alice<h2><p>I'm lucky 41</p>`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
