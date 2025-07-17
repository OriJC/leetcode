/*
 * @lc app=leetcode id=1329 lang=golang
 *
 * [1329] Sort the Matrix Diagonally
 */

// @lc code=start
func diagonalSort(mat [][]int) [][]int {
    m,n := len(mat), len(mat[0])

	diagonals := make(map[int][]int)

	for i := 0; i < m; i++{
		for j := 0; j < n; j++ {
			index := i - j
			diagonals[index] = append(diagonals[index], mat[i][j])
		}
	}

	for _, diagonal := range diagonals{
		sort.Slice(diagonal, func(i, j int) bool{
			return diagonal[i] > diagonal[j]
		})
	}

	for i := 0; i < m; i++{
		for j := 0; j < n; j++ {
			index := i - j
			mat[i][j] = diagonals[index][len(diagonals[index])-1]
			diagonals[index] = diagonals[index][:len(diagonals[index])-1]
		}
	}
	return mat
}
// @lc code=end

