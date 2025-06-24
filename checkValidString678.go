package main

func checkValidString(s string) bool {
	// 栈 + 贪心
	var leftParentheses, asterisks []int
	for i, c := range s {
		if c == '(' {
			leftParentheses = append(leftParentheses, i)
		} else if c == '*' {
			asterisks = append(asterisks, i)
		} else {
			if len(leftParentheses) > 0 {
				leftParentheses = leftParentheses[:len(leftParentheses)-1]
			} else if len(asterisks) > 0 {
				asterisks = asterisks[:len(asterisks)-1]
			} else {
				return false
			}
		}
	}

	for len(leftParentheses) > 0 && len(asterisks) > 0 {
		if leftParentheses[len(leftParentheses)-1] > asterisks[len(asterisks)-1] {
			return false
		}
		leftParentheses = leftParentheses[:len(leftParentheses)-1]
		asterisks = asterisks[:len(asterisks)-1]
	}
	return len(leftParentheses) == 0
}

//func main() {
//	// Example usage
//	s := "(*))"
//	result := checkValidString(s)
//	println(result) // Output: true
//}
