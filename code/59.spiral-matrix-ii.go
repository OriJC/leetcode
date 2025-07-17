/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
    res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	upperbound, lowerbound := 0, n - 1
	leftbound, rightbound := 0, n - 1
	num := 1
	for num <= n * n {
		if upperbound <= lowerbound{
			for j := leftbound; j <= rightbound; j++{
				res[upperbound][j] = num
				num++
			}
			upperbound++
		}

		if leftbound <= rightbound{
			for i := upperbound; i <= lowerbound; i++{
				res[i][rightbound] = num
				num++
			}
			rightbound--
		}

		if upperbound <= lowerbound{
			for j := rightbound; j >= leftbound; j--{
				res[lowerbound][j] = num
				num++
			}
			lowerbound--
		}

		if leftbound <= rightbound{
			for i := lowerbound; i >= upperbound; i--{
				res[i][leftbound] = num
				num++
			}
			leftbound++
		}

	}
	return res
}
// @lc code=end

