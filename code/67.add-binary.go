/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
    i, j := len(a)-1, len(b)-1
	carry := 0

	res := []byte{}
	for i >= 0 || j >=0 || carry != 0{
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		res = append(res, byte(sum%2) + '0')
		carry = sum / 2
	}

	for w, n := 0, len(res); w < len(res)/2 ; w++{
		res[w], res[n-1-w] = res[n-1-w], res[w]
	}
	return string(res)
}
// @lc code=end

