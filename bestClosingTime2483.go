package main

import (
	"fmt"
)

func bestClosingTime(customers string) int {
	index, min, cost := 0, 0, 0
	for i, c := range customers {
		if c == 'Y' {
			cost--
		} else {
			cost++
		}
		if cost < min {
			index = i + 1
			min = cost
		}
	}
	return index
}

func main() {
	s := "YYNY"
	// 2
	fmt.Println(bestClosingTime(s))
}
