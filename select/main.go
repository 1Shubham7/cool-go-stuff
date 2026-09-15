package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	ch2 := make( chan string)

	go func() {
		ch <- 1
	}()

	go func() {
		ch2 <- "shubham"
		close(ch2)
	}()

	select {
	case a := <- ch:
		fmt.Println(a)
	case b := <- ch2:
		fmt.Println(b) 
	}
}
