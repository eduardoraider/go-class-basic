package main

import "fmt"

func main() {

	var arr [2]string
	arr[0] = "a"
	arr[1] = "b"

	fmt.Println(arr, len(arr), cap(arr))
}
