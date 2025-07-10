/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
    n := len(nums)
	result := 0
	for i := 0; i < n-1 ; i++{
		result = max(result, i+ nums[i])
		if result <= i{
			return false
		}
	}

	return result >= n-1
}

func max(a, b int) int{
	if a > b {
		return a
	}

	return b
}
// @lc code=end

