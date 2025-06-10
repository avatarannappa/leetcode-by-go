package main

import "math"

func MaxDifference(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	maxOdd := 0
	minEven := math.MaxInt
	for _, v := range m {
		if v%2 == 0 {
			if v < minEven {
				minEven = v
			}
		} else {
			if v > maxOdd {
				maxOdd = v
			}
		}
	}
	return maxOdd - minEven
}
