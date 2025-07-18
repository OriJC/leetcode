/*
 * @lc app=leetcode id=1260 lang=golang
 *
 * [1260] Shift 2D Grid
 */

// @lc code=start
func shiftGrid(grid [][]int, k int) [][]int {
    m, n := len(grid), len(grid[0])
	mn := m*n
	k = k % mn

	// reverse the last k element
	reverse(grid, mn - k, mn - 1)
	// reverse the mn - k element
	reverse(grid, 0, mn - k - 1)
	// reverse the entire list
	reverse(grid, 0 , mn - 1)


	return grid
}

func get(grid [][]int, index int) int{
	n := len(grid[0])
	i, j := index / n, index % n
	return grid[i][j]
}

func set(grid [][]int, index int, val int){
	n := len(grid[0])
	i, j := index / n, index % n
	grid[i][j] = val
}

func reverse(grid [][]int, i, j int){
	for i < j{
		tmp := get(grid, i)
		set(grid, i, get(grid, j))
		set(grid, j, tmp)
		i++
		j--
	}
}

// @lc code=end

