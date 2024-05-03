package main

import "fmt"

func main() {
	total := add(1, 5)
	fmt.Println("total:", total)
}

func add(x, y int32) int32 {
	return x + y
}
