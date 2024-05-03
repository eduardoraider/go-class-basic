package main

import (
	"fmt"
	"time"
)

func main() {
	s := "Eduardo Raider"
	for {
		fmt.Println(s)
		time.Sleep(3 * time.Second)
	}
}
