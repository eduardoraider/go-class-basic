package main

import (
	"fmt"
	"strconv"
)

func main() {
	total, err := convertAndSumCustomError("1", "2", "5", "B")

	fmt.Println(total, err)
}

func convertAndSumCustomError(s ...string) (total int, err error) {
	for _, v := range s {
		n, e := strconv.Atoi(v)
		if e != nil {
			total = 0
			err = fmt.Errorf("the value is not a number: %s", v)
			break
		}
		total += n
	}
	return
}
