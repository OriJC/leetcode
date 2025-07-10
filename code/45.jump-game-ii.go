/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	n := len(nums)
    if n <= 1 {
		return 0
	}
	// Step jumps can jump to [i, end] 
	end, jumps := 0, 0
	farthest := 0
	for i := 0; i < n-1 ; i++{
		if nums[i] + i > farthest{
			farthest = nums[i] + i
		}

		if i == end{
			jumps++
			end = farthest
			if farthest >= n-1{
				return jumps
			}
		}
	}
	return -1
}
// @lc code=end

