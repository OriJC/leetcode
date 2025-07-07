/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minDepth := math.MaxInt32

	currentDepth := 0

	var traverse func(*TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		currentDepth++
		if root.Left == nil && root.Right == nil {
			minDepth = min(minDepth, currentDepth)
		}
		traverse(root.Left)
		traverse(root.Right)
		currentDepth--
	}

	traverse(root)
	return minDepth
}

// @lc code=end

