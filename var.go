package main

import "fmt"

var (
	age uint8  = 47 // var global scope, cant not use := in global scope
	doc string = "000.000.000-00"
)

func main() {
	name := "Eduardo Raider" // var function scope

	var (
		address string = "Street A"
		city    string = "Goland"
	)

	fmt.Println(name, age, doc, address, city)
}
