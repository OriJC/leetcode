/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start
func nthUglyNumber(n int) int {
    ugly := make([]int, n)
	ugly[0] = 1


	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n ; i++{
		next2 := ugly[i2] * 2
		next3 := ugly[i3] * 3
		next5 := ugly[i5] * 5

		next := min(next2, min(next3, next5))
		ugly[i] = next

		if next == next2{
			i2++
		}
		if next == next3{
			i3++
		}
		if next == next5{
			i5++
		}
	}
	return ugly[n-1]
}

func min(a, b int) int{
	if a < b{
		return a
	}

	return b
}
// @lc code=end

