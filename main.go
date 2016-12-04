package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/skiarn/browse-shows/templator"
	"github.com/skiarn/browse-shows/types/show"
)

func main() {
	var port = flag.Int("p", 8080, "What port the mock should run on.")
	var url = flag.String("u", "", "api endpoint")
	flag.Parse() // parse the flags

	showtmpl := templator.New(show.RemoteAPIURL(url), show.Render(), "^/(view)/([a-zA-Z0-9_-]+)$", "view.html")
	http.HandleFunc("/view/", templator.Handler(templator.ViewHandler, showtmpl))

	log.Printf("Server to listen on a port: %v \n", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
