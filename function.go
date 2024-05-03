package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		hello()
	}
}

func hello() {
	fmt.Println("Hello Eduardo")
}
