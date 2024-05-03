package main

import "fmt"

func main() {
	age := 8
	height := 90

	switch {
	case age >= 6:
		fmt.Println("can play")
	case age >= 4 && height > 100:
		fmt.Println("can play together to")
	default:
		fmt.Println("can't play")
	}
}
