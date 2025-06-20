package main

import "math"

var str string

func maxDistance(s string, k int) int {
	// 贪心 + 枚举
	str = s
	max := 0.0

	max = math.Max(max, cal('N', 'E', k))
	max = math.Max(max, cal('S', 'E', k))
	max = math.Max(max, cal('S', 'W', k))
	max = math.Max(max, cal('N', 'W', k))

	return int(max)
}

func cal(a rune, b rune, k int) float64 {
	max := 0.0
	steps := 0.0
	for _, v := range str {
		if v == a || v == b {
			steps++
		} else {
			if k > 0 {
				k--
				steps++
			} else {
				steps--
			}
		}
		max = math.Max(max, steps)
	}
	return max
}

//func main() {
//	s := "NWSE"
//	k := 1
//	result := maxDistance(s, k)
//	println(result) // Output the result
//}
