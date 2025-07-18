/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
    result := [][]int {}
    
    n := len(nums)

    if n < 4 {
        return result
    }

    sort.Ints(nums)
    
    for i := 0; i < n - 3; i++ {
        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        for j := i + 1; j < n - 2; j++{
            if j > i+1 && nums[j] == nums[j-1]{
                continue
            }


            left, right := j+1, n-1

            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]

                if sum == target {
                    result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

                    for left < right && nums[left] == nums[left+1]{
                        left++
                    }

                    for left < right && nums[right] == nums[right-1]{
                        right--
                    }
                    left ++
                    right --
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }
        }
    }
    return result
}
// @lc code=end

