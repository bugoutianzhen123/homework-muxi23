package main

import (
	"fmt"
	"time"
)

func printword(c1, c2 chan int) {
	for i := 65; i < 91; i += 2 {
		<-c1
		fmt.Printf("%c%c", i, i+1)
		c2 <- 0
	}
}

func printnum(c1, c2 chan int) {
	for i := 0; i < 26; i += 2 {
		<-c2
		fmt.Printf("%d%d", i, i+1)
		c1 <- 0
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go printnum(ch1, ch2)
	go printword(ch1, ch2)
	ch1 <- 0
	time.Sleep(1 * time.Second)
}
