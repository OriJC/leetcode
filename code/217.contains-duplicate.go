/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {

	count := make(map[int]bool)
	for _, num := range nums {
		if count[num] {
			return true
		}
		count[num] = true
	}
	return false
}

// @lc code=end

