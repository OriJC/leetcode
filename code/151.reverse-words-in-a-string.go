/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1{
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
// @lc code=end

