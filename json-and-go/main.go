package main

import (
	"encoding/json"
	"fmt"
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

func printInterfaceData(i interface{}) {

	mi := i.(map[string]interface{})
	log.Printf("%s", mi)

	for k, v := range mi {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

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
	var i interface{}
	err_in := json.Unmarshal(b, &i)
	if err_in == nil {
		log.Printf("%s", i)
	}

	printInterfaceData(i)
}
