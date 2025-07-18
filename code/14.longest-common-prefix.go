/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
    m, n := len(strs), len(strs[0])
	for col := 0; col < n; col++{
		for row := 1; row < m; row++{
			thisStr, prevStr := strs[row], strs[row-1]
			if col >= len(thisStr) || col >= len(prevStr) || thisStr[col] != prevStr[col]{
				return strs[row][:col]
			}
		}
	}
	return strs[0]
}
// @lc code=end

