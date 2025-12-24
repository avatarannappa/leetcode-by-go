package main

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	sum := 0
	for _, i := range apple {
		sum += i
	}
	result := 0
	for i := range capacity {
		if capacity[i] >= sum {
			return result + 1
		}
		result += 1
		sum -= capacity[i]
	}
	return -1
}
