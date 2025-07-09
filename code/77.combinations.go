/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
    var res [][]int

	var track []int

	var backtrack func(int)
	backtrack = func(start int){
		if k == len(track){
			res = append(res, append([]int(nil), track...))
			return
		}

		for i := start ; i <= n; i++ {
			track = append(track, i)

			backtrack(i+1)
			
			track = track[:len(track)-1]
		}
	} 

	backtrack(1)
	return res
}
// @lc code=end

