package main

import "fmt"

func main() {
	arr := [2]string{"Eduardo", "Raider"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
