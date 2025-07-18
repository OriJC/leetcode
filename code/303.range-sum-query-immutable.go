/*
 * @lc app=leetcode id=303 lang=golang
 *
 * [303] Range Sum Query - Immutable
 */

// @lc code=start
type NumArray struct {
    PreSum []int
}


func Constructor(nums []int) NumArray {
    preSum := make([]int, len(nums)+1)

	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	return NumArray{PreSum: preSum}
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.PreSum[right+1] - this.PreSum[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
// @lc code=end

