/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range s1 {
		need[c]++
	}

	valid := 0
	left, right := 0, 0

	for right < len(s2) {
		c := rune(s2[right])
		right++
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := rune(s2[left])
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

// @lc code=end

