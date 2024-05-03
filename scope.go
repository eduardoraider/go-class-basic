package main

import "fmt"

var name string = "Eduardo Raider"

func main() {
	// name := "John Doe"
	// fmt.Println(name)

	var name string
	if len(name) == 0 {
		name = "Jane Doe"
	}

	fmt.Printf("Hello, %s!\n", name)
	fmt.Println("======")
	show()
}

func show() {
	fmt.Println(name)
}
