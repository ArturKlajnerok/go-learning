package main

import (
	"github.com/mailgun/oxy/forward"
	"github.com/mailgun/oxy/testutils"
	"net/http"
)

func main() {
	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	fwd, _ := forward.New()

	redirect := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// let us forward this request to another server
		req.URL = testutils.ParseURI("http://www.google.com")
		fwd.ServeHTTP(w, req)
	})

	// that's it! our reverse proxy is ready!
	s := &http.Server{
		Addr:    ":8080",
		Handler: redirect,
	}
	s.ListenAndServe()
}
