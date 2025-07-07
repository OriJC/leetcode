/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range p {
		need[c]++
	}

	left, right, valid := 0, 0, 0
	res := []int{}
	for right < len(s) {
		c := rune(s[right])
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}
			d := rune(s[left])
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

// @lc code=end

