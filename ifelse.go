package main

import "fmt"

func main() {

	animal := "dog"

	if animal == "cat" {
		fmt.Println("meow")
	} else if animal == "dog" {
		fmt.Println("au au")
	} else if animal == "bird" {
		fmt.Println("piu piu")
	} else {
		fmt.Println("unknown animal")
	}
}
