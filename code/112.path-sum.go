/*
 * @lc app=leetcode id=112 lang=golang
 *
 * [112] Path Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return EqualTarget(root, 0, targetSum)
}

func EqualTarget(root *TreeNode, currSum int, targetSum int) bool {
	if root == nil {
		return false
	}

	currSum += root.Val
	if root.Left == nil && root.Right == nil {
		return currSum == targetSum
	}
	return EqualTarget(root.Left, currSum, targetSum) || EqualTarget(root.Right, currSum, targetSum)
}

// @lc code=end

