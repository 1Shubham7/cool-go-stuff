package main

import "fmt"

type animal interface {
	Bark(string) string
}

type Dog struct {
	Name string
	age int
}

func (d *Dog) Bark(input string) string {
	return fmt.Sprintf("dog is barking: %s", input)
}

func main() {
	var a animal = &Dog{
		Name: "shubham",
		age: 12,
	}

	fmt.Println(a.Bark("woff woff"))
}