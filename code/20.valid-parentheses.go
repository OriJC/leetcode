/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
func isValid(s string) bool {
	left := []rune{}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			left = append(left, char)

		} else {
			if len(left) > 0 && leftOf(char) == left[len(left)-1] {
				left = left[:len(left)-1]
			} else {
				return false
			}
		}
	}
	return len(left) == 0
}

func leftOf(c rune) rune {
	if c == ')' {
		return '('
	} else if c == '}' {
		return '{'
	} else {
		return '['
	}
}

// @lc code=end

