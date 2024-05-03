package main

import "fmt"

func main() {

	animal := "bear"

	switch animal {
	case "cat":
		fmt.Println("meow")
	case "dog":
		fmt.Println("au au")
	case "bird":
		fmt.Println("piu piu")
	default:
		fmt.Println("unknown animal")
	}

}
