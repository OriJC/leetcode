/*
 * @lc app=leetcode id=773 lang=golang
 *
 * [773] Sliding Puzzle
 */

// @lc code=start
func slidingPuzzle(board [][]int) int {
	target := "123450"

	start := ""

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			start += string(rune(board[i][j]) + '0')
		}
	}
	step := 0
	queue := []string{start}
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur == target {
				return step
			}

			for _, neighbor := range getNeighbors(cur) {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
			}
		}
		step++
	}
	return -1
}

func getNeighbors(board string) []string {
	mapping := [][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}

	idx := strings.Index(board, "0")
	neighbors := []string{}
	for _, adj := range mapping[idx] {
		newBoard := swap(board, idx, adj)
		neighbors = append(neighbors, newBoard)
	}

	return neighbors
}

func swap(board string, i, j int) string {
	chars := []rune(board)
	chars[i], chars[j] = chars[j], chars[i]
	return string(chars)
}

// @lc code=end

