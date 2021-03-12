package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	var chan1 = make(chan string)
	var chan2 = make(chan string)
	go func() {
		for {
			chan1 <- "hello"
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for {
			chan2 <- "world"
			time.Sleep(2 * time.Second)
		}
	}()
	for {
		select {
		case message1 := <-chan1:
			fmt.Println(message1)
		case message2 := <-chan2:
			fmt.Println(message2)
		}
	}
}
