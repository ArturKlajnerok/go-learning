package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

	// read directory
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// read directory recursively
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
