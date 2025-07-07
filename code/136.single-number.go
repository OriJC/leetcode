/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	for _, num := range nums {
		if count[num] == 1 {
			return num
		}
	}

	return -1
}

// @lc code=end

