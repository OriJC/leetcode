/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	window, need := make(map[byte]int), make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	start, length := 0, math.MaxInt32
	valid := 0
	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}

		}
	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length]
}

// @lc code=end

