package main

import (
	"fmt"
	"runtime"
)

func printword() {
	for i := 65; i < 91; i += 2 {
		fmt.Printf("%c%c", i, i+1)
		runtime.Gosched()
	}
}

func printnum() {
	for i := 0; i < 26; i += 2 {
		fmt.Printf("%d%d", i, i+1)
		runtime.Gosched()
	}
}

func main() {
	go printnum()
	printword()
}
