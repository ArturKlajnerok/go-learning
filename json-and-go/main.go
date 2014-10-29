package main

import (
	"encoding/json"
	"log"
	//"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

var m Message
var mu Message

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	b, err_ma := json.Marshal(m)
	if err_ma == nil {
		log.Printf("%s", b)
		//os.Stdout.Write(b)
	}

	err_um := json.Unmarshal(b, &mu)

	if err_um == nil {
		log.Printf("%s", mu.Name)
	}
}
