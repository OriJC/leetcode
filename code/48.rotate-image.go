/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int)  {
    n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range(matrix){
		reverse(row)
	}
}

func reverse(arr []int){
	i, j := 0, len(arr)-1
	for j > i {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
// @lc code=end

