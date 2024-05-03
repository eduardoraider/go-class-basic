package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Doc  string
	Addr string
	City string
}

func main() {

	p := Person{
		Name: "Eduardo Raider",
		Age:  47,
		Doc:  "000.000.000-00",
		Addr: "Street One",
		City: "Goland",
	}

	fmt.Println(p)
}
