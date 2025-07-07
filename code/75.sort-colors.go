/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	max, min := math.MinInt64, math.MaxInt64
	for _, num := range nums {
		min = minInt(min, num)
		max = maxInt(max, num)
	}

	offset := -min
	count := make([]int, max-min+1)

	for _, num := range nums {
		count[num+offset]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	sorted := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		sorted[count[nums[i]+offset]-1] = nums[i]
		count[nums[i]+offset]--
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = sorted[i]
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

