package main

import "fmt"

func main() {
	n := []uint8{1, 2, 3, 4, 5}

	// descending order
	for i := len(n) - 1; i >= 0; i-- {
		fmt.Println(n[i])
	}

	fmt.Println("==========")

	// ascending order
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i])
	}

	fmt.Println("==========")

	// "foreach" slice - ascending order
	for k, v := range n {
		fmt.Println(k, v)
	}

}
