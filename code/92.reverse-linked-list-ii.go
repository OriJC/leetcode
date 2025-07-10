/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    var successor *ListNode

	var reverseN func(head *ListNode, n int) *ListNode
	reverseN = func(head *ListNode, n int) *ListNode{
		if n == 1 {
			successor = head.Next
			return head
		}

		last := reverseN(head.Next, n-1)

		head.Next.Next = head
		head.Next = successor
		return last
	}

	if m == 1{
		return reverseN(head, n)
	}

	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}
// @lc code=end

