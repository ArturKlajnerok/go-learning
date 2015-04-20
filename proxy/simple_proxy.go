package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"unicode"
)

var (
	listen = flag.String("listen", "localhost:1080", "listen on address")
	logp   = flag.Bool("log", true, "enable logging")
)

func main() {
	flag.Parse()
	proxyHandler := http.HandlerFunc(proxyHandlerFunc)
	log.Fatal(http.ListenAndServe(*listen, proxyHandler))
}

func proxyHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// Log if requested
	if *logp {
		log.Println(r.URL)
		log.Println(r.RequestURI)
	}

	// We'll want to use a new client for every request.
	client := &http.Client{}

	// Tweak the request as appropriate:
	// RequestURI may not be sent to client
	// URL.Scheme must be lower-case
	r.RequestURI = ""
	r.URL.Scheme = strings.Map(unicode.ToLower, r.URL.Scheme)

	// And proxy
	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	w.Write(body)
}
