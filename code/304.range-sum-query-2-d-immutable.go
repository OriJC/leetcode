/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start
type NumMatrix struct {
    PreSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0{
		return NumMatrix{}
	}
	preSum := make([][]int, m+1)
	for i := range preSum{
		preSum[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] + matrix[i-1][j-1] - preSum[i-1][j-1]
		}
	}
	return NumMatrix{PreSum: preSum}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.PreSum[row2+1][col2+1] - this.PreSum[row2+1][col1] - this.PreSum[row1][col2+1] + this.PreSum[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

