/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
    var res [][]int
	var track []int

	sort.Ints(nums)
	backtrack(nums, &track, &res, 0)
	return res
}

func backtrack(nums []int, track *[]int, res *[][]int, start int){
	tmp := make([]int, len(*track))
	copy(tmp, *track)
	*res = append(*res, tmp)

	for i := start; i < len(nums); i++{
		if i > start && nums[i] == nums[i-1]{
			continue
		}

		*track = append(*track, nums[i])
		backtrack(nums, track, res, i+1)
		*track = (*track)[:len(*track)-1]
	}

}

// @lc code=end

