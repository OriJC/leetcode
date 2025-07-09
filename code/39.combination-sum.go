/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
	var track []int
	tracksum := 0

	var backtrack func([]int, int)

	backtrack = func(nums []int, start int){
		if tracksum == target {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		if tracksum > target{
			return
		}

		for i := start; i < len(nums); i++{
			tracksum += nums[i]
			track = append(track, nums[i])

			backtrack(nums, i)

			tracksum -= nums[i]
			track = track[:len(track)-1]
		}
	}

	backtrack(candidates, 0)
	return res
}

// @lc code=end

