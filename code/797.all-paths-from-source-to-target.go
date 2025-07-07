/*
 * @lc app=leetcode id=797 lang=golang
 *
 * [797] All Paths From Source to Target
 */

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var path []int
	traverse(graph, 0, &path, &res)
	return res
}

func traverse(graph [][]int, s int, path *[]int, res *[][]int) {
	*path = append(*path, s)

	n := len(graph)

	if s == n-1 {
		*res = append(*res, append([]int(nil), *path...))
		*path = (*path)[:len(*path)-1]
		return
	}

	for _, v := range graph[s] {
		traverse(graph, v, path, res)
	}

	*path = (*path)[:len(*path)-1]
}

// @lc code=end

