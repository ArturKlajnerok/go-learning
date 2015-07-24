package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"http://google.ac/",
	"http://google.ad/",
	"http://google.ae/",
	"http://google.al/",
	"http://google.am/",
	"http://google.as/",
	"http://google.at/",
	"http://google.az/",
	"http://google.ba/",
	"http://google.be/",
	"http://google.bf/",
	"http://google.bg/",
	"http://google.bi/",
	"http://google.bj/",
	"http://google.bs/",
	"http://google.bt/",
	"http://google.by/",
	"http://google.com/",
	"http://google.com/404",
	"http://google.nohost/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string) <-chan *HttpResponse {
	ch := make(chan *HttpResponse)
	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			if resp != nil {
				resp.Body.Close()
			}
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}

	return ch
}

func main() {
	results := asyncHttpGets(urls)
	fmt.Println("Results :")
	for _ = range urls {
		result := <-results
		if result.response != nil {
			fmt.Printf("%s status: %s\n", result.url, result.response.Status)
		} else if result.err != nil {
			fmt.Printf("%s error: %s\n", result.url, result.err)
		}
	}
}
