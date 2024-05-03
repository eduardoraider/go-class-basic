package main

import "fmt"

func main() {
	s := "Eduardo Raider"
	for i := 0; i < len(s); i++ {
		// fmt.Println(s[i], string(s[i]))
		// fmt.Println(s[i], fmt.Sprintf("%c", s[i]))
		fmt.Println(s[i], fmt.Sprintf("%s", string(s[i])))
	}
}
