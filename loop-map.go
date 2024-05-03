package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "John",
		"course": "golang",
		"site":   "imooc",
	}

	for k, v := range m {
		println(k, v)
	}

	fmt.Println("==========")

	// avoid in map
	m2 := map[int]string{
		0: "one",
		1: "two",
		2: "three",
	}

	for i := 0; i <= len(m2); i++ {
		println(m2[i])
	}
}
