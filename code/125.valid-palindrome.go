/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
func isPalindrome(s string) bool {
    var sb strings.Builder
	for i := 0; i < len(s); i++{
		c := s[i]
		if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)){
			sb.WriteByte(byte(unicode.ToLower(rune(c))))
		}
	}

	filtered := sb.String()
	left, right := 0, len(filtered) - 1
	for left < right{
		if filtered[left] != filtered[right]{
			return false
		}
		left++
		right--
	}

	return true
}
// @lc code=end

