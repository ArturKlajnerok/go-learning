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

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

var m_new Message
var i interface{}

func encodeMessage(m Message) []byte {
	log.Print("# Encoding Message")
	b, err := json.Marshal(m)
	if err == nil {
		log.Printf("%s", b)
		//os.Stdout.Write(b)
	}
	return b
}

func decodeMessageStruct(b []byte) {
	log.Print("# Decoding Message to struct")
	err := json.Unmarshal(b, &m_new)
	if err == nil {
		log.Printf("%s", m_new.Name)
	}
}

func decodeMessageInterface(b []byte) {
	log.Print("# Decoding Message to interface")
	err := json.Unmarshal(b, &i)
	if err == nil {
		log.Printf("%s", i)
		printInterfaceData(i)
	}
}

func decodeFamilyMemberInterface(b []byte) {
	log.Print("# Decoding FamilyMember to interface")
	//var fm FamilyMember
	//err_fm := json.Unmarshal(by, &fm)
	err := json.Unmarshal(b, &i)
	if err == nil {
		//log.Printf("%s", fm)
		log.Printf("%s", i)
		printInterfaceData(i)
	}
}

func printInterfaceData(i interface{}) {
	log.Print("# Prinitng interface data")

	ma := i.(map[string]interface{})
	for k, v := range ma {
		switch vv := v.(type) {
		case string:
			log.Print(k, " is string ", vv)
		case int:
			log.Print(k, " is int ", vv)
		case []interface{}:
			log.Print(k, " is an array: ")
			for i, u := range vv {
				log.Print(i, " ", u)
			}
		default:
			log.Print(k, " is of a type I don't know how to handle")
		}
	}
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	bm := encodeMessage(m)

	decodeMessageStruct(bm)
	decodeMessageInterface(bm)

	bfm := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	decodeFamilyMemberInterface(bfm)

}
