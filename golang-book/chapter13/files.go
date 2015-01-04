package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)

	// shorter read fle
	bs2, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str2 := string(bs2)
	fmt.Println(str2)

	// create file
	file2, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file2.Close()

	file2.WriteString("Go new test")
}
