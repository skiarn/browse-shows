package program 

import (
	"html/template"
	"testing"
	"net/http"
	"net/http/httptest"
	)	

func TestVisitURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/view/idol", nil)
	if err != nil {
		t.Fatal(err)
	}

	testFunc := func(w http.ResponseWriter, r *http.Request, title string){
		p := &Program{Title: title}
		renderTemplate(w, "template test", p)
		return
	} 	

	rr := httptest.NewRecorder()
	//set templates to test templates
	testTemplates := templates
	//we are exiting, revert back templates to testTemplates at end of function.
	defer func() { templates = testTemplates }()
	ttemplate := template.New("template test.html")
    	templates = template.Must(ttemplate.Parse("template test {{.Title}}"))
	handler := http.HandlerFunc(Handler(testFunc))

    	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	
	expected := `template test idol`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", 
			rr.Body.String(), expected)
	}	
}
