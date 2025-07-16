package main

func maximumLength(nums []int) int {
	// 贪心 枚举四种情况：全偶，全奇，奇偶，偶奇
	a := 0
	b := 0
	c := 0
	d := 0

	ce := false
	de := true

	for _, v := range nums {
		if v%2 == 0 {
			a++
			if !ce {
				c++
				ce = !ce
			}
			if !de {
				d++
				de = !de
			}
		} else {
			b++
			if ce {
				c++
				ce = !ce
			}
			if de {
				d++
				de = !de
			}
		}
	}

	return max(a, b, c, d)
}

func max(a, b, c, d int) int {
	if a < b {
		a = b
	}
	if a < c {
		a = c
	}
	if a < d {
		a = d
	}
	return a
}
