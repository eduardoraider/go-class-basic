package main

import (
	"fmt"
	"strconv"
)

func main() {
	total, err := convertAndSum("1", "2", "5")
	// total, err := convertAndSum("1", "2", "5", "B")

	fmt.Println(total, err)
}

func convertAndSum(s ...string) (total int, err error) {
	var n int
	for _, v := range s {
		n, err = strconv.Atoi(v)
		if err != nil {
			total = 0
			break
		}
		total += n
	}
	return
}
