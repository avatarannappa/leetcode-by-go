package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Sort(sort.Reverse(sort.IntSlice(happiness)))
	var ans int64
	j := 0
	for i := range happiness {
		if k == 0 {
			return ans
		}
		k -= 1
		num := happiness[i]
		if num > j {
			ans += int64(num - j)
		} else {
			return ans
		}
		j++
	}
	return ans
}

func maximumHappinessSumNew(happiness []int, k int) int64 {
	sort.Slice(happiness, func(a, b int) bool { return b-a > 0 })
	var ans int64
	for i, x := range happiness[:k] {
		if x <= i {
			break
		}
		ans += int64(x - i)
	}
	return ans
}
