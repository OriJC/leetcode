/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int)
	start := 0
	length := 0
	for i, c := range s {
		if idx, found := charIndex[c]; found && idx >= start {
			start = idx + 1
		}
		charIndex[c] = i
		if i-start+1 > length {
			length = i - start + 1
		}
	}
	return length
}

// @lc code=end

