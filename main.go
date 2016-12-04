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
	var link = flag.String("link", "", "web link")

	flag.Parse() // parse the flags

	showtmpl := templator.New(show.RemoteAPIURL(url), show.Render(), "^/(view)/([a-zA-Z0-9_-]+)$", "view.html")
	http.HandleFunc("/view/", templator.Handler(templator.ViewHandler, showtmpl))

	participanttmpl := templator.New(show.RemoteParticipantAPIURL(url), show.RenderParticipant(link), "^/(view)/participant/([a-zA-Z_-]+)$", "view-participant.html")
	http.HandleFunc("/view/participant/", templator.Handler(templator.ViewHandler, participanttmpl))

	log.Printf("Server to listen on a port: %v \n", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
