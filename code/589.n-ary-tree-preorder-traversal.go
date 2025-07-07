/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, child := range root.Children {
		res = append(res, preorder(child)...)
	}
	return res
}

// @lc code=end

