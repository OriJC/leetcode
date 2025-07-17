/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
    slow, fast := 0, 0
	count := 0
	for fast < len(nums){
		if nums[fast] != nums[slow]{
			slow++
			nums[slow] = nums[fast]
		} else if slow < fast && count < 2{
			slow++
			nums[slow] = nums[fast]
		}

		count++
		fast++
		if fast < len(nums) && nums[fast]!=nums[fast-1]{
			count = 0
		}

	}
	return slow + 1 
}
// @lc code=end

