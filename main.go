package main

import (
	"flag"
	"log"
	"net/http"
	"program"
	"strconv"
)

func main() {
	var port = flag.Int("p", 8080, "What port the mock should run on.")
	var url = flag.String("u", "", "api endpoint")
	flag.Parse() // parse the flags
	http.HandleFunc("/view/", program.Handler(program.ViewHandler, program.New(*url, "^/(view)/([a-zA-Z0-9_-]+)$", "view.html")))

	log.Printf("Server to listen on a port: %v \n", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
