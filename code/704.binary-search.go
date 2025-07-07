/*
 * @lc app=leetcode id=704 lang=golang
 *
 * [704] Binary Search
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	index := len(nums) / 2

	if nums[index] == target {
		return index
	} else if nums[index] < target {
		result := search(nums[index+1:], target)
		if result == -1 {
			return -1
		}
		return result + index + 1
	} else {
		return search(nums[:index], target)
	}
}

// @lc code=end

