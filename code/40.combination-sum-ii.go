/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int
	var track []int
	var tracksum int

	sort.Ints(candidates)
	backtrack(candidates, 0, target, &track, &res, &tracksum)
	return res
}

func backtrack(nums []int, start, target int, track *[]int, res*[][]int, tracksum *int){
	if *tracksum == target {
		tmp := make([]int, len(*track))
		copy(tmp, *track)
		*res = append(*res, tmp)
		return
	}

	if *tracksum > target{
		return
	}

	for i := start; i < len(nums); i++{
		if i > start && nums[i] == nums[i-1]{
			continue
		}

		*track = append(*track, nums[i])
		*tracksum += nums[i]

		backtrack(nums, i+1, target, track, res, tracksum)

		*track = (*track)[:len(*track)-1]
		*tracksum -= nums[i]
	}
}
// @lc code=end

