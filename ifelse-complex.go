package main

import "fmt"

func main() {
	age := 4
	height := 90

	if age >= 6 {
		fmt.Println("can play")
	} else if age >= 4 && height > 100 {
		fmt.Println("can play together to")
	} else {
		fmt.Println("can't play")
	}
}
