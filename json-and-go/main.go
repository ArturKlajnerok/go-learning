package main

import (
	"encoding/json"
	//"fmt"
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
var mi Message

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	// Encode
	b, err_ma := json.Marshal(m)
	if err_ma == nil {
		log.Printf("%s", b)
		//os.Stdout.Write(b)
	}

	// Decode to struct
	err_um := json.Unmarshal(b, &mu)

	if err_um == nil {
		log.Printf("%s", mu.Name)
	}

	// Decoding to interface
	var inter interface{}
	err_in := json.Unmarshal(b, &inter)
	if err_in == nil {
		log.Printf("%s", inter)
	}

}
