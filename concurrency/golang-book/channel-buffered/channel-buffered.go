package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		fmt.Println(i)
		c <- "ping"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string, 10)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
