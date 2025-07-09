/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
    var res [][]int
	var track []int

	used := make([]bool, len(nums))

	sort.Ints(nums)
	backtrack(nums, &res, &track, used)
	return res

}

func backtrack(nums []int, res *[][]int, track *[]int, used []bool){
	if len(*track) == len(nums){
		tmp := make([]int, len(*track))
		copy(tmp, *track)
		*res = append(*res, tmp)
		return 
	}

	for i := 0; i < len(nums); i++{
		if used[i]{
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1]{
			continue
		}

		*track = append(*track, nums[i])
		used[i] = true

		backtrack(nums, res, track, used)

		*track = (*track)[:len(*track)-1]
		used[i] = false
	}

}
// @lc code=end

