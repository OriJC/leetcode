/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
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
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var traverse func(root *TreeNode)

	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		traverse(root.Left)
		traverse(root.Right)
	}

	traverse(root)
	return res
}

// @lc code=end

