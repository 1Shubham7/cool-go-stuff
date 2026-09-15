package main

import "fmt"


func main() { 
	var a any = "shubham"

	switch a.(type) {
	case string: 
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("none of the above")
	}
}