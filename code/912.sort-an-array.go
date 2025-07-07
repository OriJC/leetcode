/*
 * @lc app=leetcode id=912 lang=golang
 *
 * [912] Sort an Array
 */

// @lc code=start
func sortArray(nums []int) []int {
	max, min := math.MinInt64, math.MaxInt64
	for _, num := range nums {
		max = maxInt(max, num)
		min = minInt(min, num)
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

	return sorted
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

