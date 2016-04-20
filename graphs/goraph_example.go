package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gyuho/goraph"
)

func ExampleNewGraph() {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	g, err := goraph.NewGraphFromJSON(f, "graph_00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.String())
}
