/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
    var res []int
	m := len(matrix)
	n := len(matrix[0])
	upperbound := 0
	lowerbound := m - 1
	leftbound := 0
	rightbound := n - 1
	for len(res) < m * n{
		if upperbound <= lowerbound{
			for j := leftbound; j <= rightbound ; j++{
				res = append(res, matrix[upperbound][j])
			}
			upperbound++
		}

		if leftbound <= rightbound {
			for i := upperbound; i <= lowerbound; i++{
				res = append(res, matrix[i][rightbound])
			}
			rightbound--
		}

		if upperbound <= lowerbound{
			for j := rightbound; j >= leftbound ; j--{
				res = append(res, matrix[lowerbound][j])
			}
			lowerbound--
		}

		if leftbound <= rightbound {
			for i := lowerbound; i >= upperbound; i--{
				res = append(res, matrix[i][leftbound])
			}
			leftbound++
		}
	}
	return res
}
// @lc code=end

