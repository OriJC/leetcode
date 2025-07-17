/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	n := len(nums)
    i, j := 0, n-1
	res := make([]int, n)
	p := n-1
	for i <= j{
		if abs(nums[i]) > abs(nums[j]){
			res[p] = nums[i] * nums[i]
			i++
		} else {
			res[p] = nums[j] * nums[j]
			j--
		}
		p--
	}
	return res
}

func abs(a int) int{
	if a < 0{
		return -a
	}
	return a
}
// @lc code=end

