package main

import "fmt"

type Person2 struct {
	Name string
	Age  int
	Doc  string
	Addr string
	City string
}

func main() {

	p := &Person2{
		Name: "Eduardo Raider",
	}

	fmt.Println(p)

	p2 := new(Person2)
	p2.Name = "Eduardo Raider"

	fmt.Println(p2)
}
