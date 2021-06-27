package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vzms/content"
)

// NOTE: package-level declarations okay here as part of startup wiring,
// as long as functionality lives in other packages that do get wired correctly
// without package-level/global references

var siteRouter *httprouter.Router
var contentHandler *content.Handler

func main() {

	httpF := flag.String("http", ":9080", "HTTP host:port to listen on for requests")
	flag.Parse()

	// TODO: possibly admin and login should just be disabled in production, for this setup, add flag(s) for it

	// NOTES:
	// server-side vugu pages
	// admin wasm
	// content API
	// login

	siteRouter = httprouter.New()
	contentHandler = &content.Handler{}

	addRoutes()

	log.Printf("Starting HTTP server at %q", *httpF)
	s := &http.Server{
		Addr:    *httpF,
		Handler: siteRouter,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
