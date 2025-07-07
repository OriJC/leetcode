/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	deads := make(map[string]struct{})
	for _, dead := range deadends {
		deads[dead] = struct{}{}
	}

	if _, found := deads["0000"]; found {
		return -1
	}

	visited := make(map[string]struct{})

	q := make([]string, 0)

	step := 0
	q = append(q, "0000")
	visited["0000"] = struct{}{}

	for len(q) > 0 {
		sq := len(q)
		for i := 0; i < sq; i++ {
			cur := q[0]
			q = q[1:]

			if cur == target {
				return step
			}

			for _, neighbor := range getNeighbors(cur) {
				if _, found := visited[neighbor]; !found {
					if _, dead := deads[neighbor]; !dead {
						q = append(q, neighbor)
						visited[neighbor] = struct{}{}
					}
				}
			}

		}
		step++
	}
	return -1
}

func plusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j]++
	}
	return string(ch)
}

func minusOne(s string, j int) string {
	ch := []rune(s)
	if ch[j] == '0' {
		ch[j] = '9'
	} else {
		ch[j]--
	}
	return string(ch)
}

func getNeighbors(s string) []string {
	neighbors := make([]string, 0)
	for i := 0; i < 4; i++ {
		neighbors = append(neighbors, plusOne(s, i))
		neighbors = append(neighbors, minusOne(s, i))
	}
	return neighbors
}

// @lc code=end

