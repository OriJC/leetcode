/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMax := traverse(root.Left)
		rightMax := traverse(root.Right)

		maxDiameter = max(maxDiameter, leftMax+rightMax)
		return 1 + max(leftMax, rightMax)
	}
	traverse(root)
	return maxDiameter
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

