/*
 * @lc app=leetcode id=365 lang=golang
 *
 * [365] Water and Jug Problem
 */

// @lc code=start
func canMeasureWater(x int, y int, target int) bool {
	if target == 0 {
		return true
	}

	if x+y < target {
		return false
	}

	return target%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// @lc code=end

