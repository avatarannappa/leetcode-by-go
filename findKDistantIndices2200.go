package main

func findKDistantIndices(nums []int, key int, k int) []int {
	var result []int
	var indices []int
	for i, v := range nums {
		if v == key {
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		return result
	}

	j := 0
	for i := 0; i < len(nums); {
		if i < indices[j] {
			if indices[j]-i <= k {
				result = append(result, i)
			}
			i++
		} else {
			if i-indices[j] <= k {
				result = append(result, i)
				i++
			} else {
				j++
				if j >= len(indices) {
					break
				}
			}
		}
	}
	return result
}

//func main() {
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	key := 2
//	k := 5
//	// [0, 6]
//	result := findKDistantIndices(nums, key, k)
//	fmt.Println(result)
//}
