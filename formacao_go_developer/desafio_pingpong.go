package main

import (
	"fmt"
	"time"
)

func ping (c1 chan string) {

	for i := 0; ; i++ {
		c1 <- "ping"
	}
}

func pong (c2 chan string) {

	for i := 0; ; i++ {
		c2 <- "pong"
	}
}

func imprimir (c1, c2 chan string) {

	for {
		p1 := <- c1
		fmt.Println(p1)
		time.Sleep(1 * time.Second)
		p2 := <- c2
		fmt.Println(p2)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)

	go ping(c1)
	go pong(c2)
	go imprimir(c1, c2)

	var entrada string
	fmt.Scanln(&entrada)
}