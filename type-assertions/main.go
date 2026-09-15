package main

import "fmt"



func main() {
	var a any = 1
	
	value, ok := a.(int)
	if ok {
		fmt.Println("a is int", value)
	}
}
