/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}

	track := []int{}
	used := make([]bool, len(nums))

	backtrack(nums, track, used, &res)
	return res
}

func backtrack(nums []int, track []int, used []bool, res *[][]int) {

	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}

		track = append(track, nums[i])
		used[i] = true
		backtrack(nums, track, used, res)
		track = track[:len(track)-1]
		used[i] = false
	}
}

// @lc code=end

