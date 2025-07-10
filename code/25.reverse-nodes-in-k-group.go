/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
		return head
	}

	a, b := head, head
	for i := 0 ; i < k ; i++{
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverseN(a, k)

	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverseN(head *ListNode, n int) *ListNode{
	if head == nil || head.Next == nil {
		return head
	}

	pre, cur, nxt := (*ListNode)(nil), head, head.Next
	for n > 0 {
		cur.Next = pre
		pre = cur
		cur = nxt
		if nxt != nil{
			nxt = nxt.Next
		}
		n--
	}

	head.Next = cur
	return pre
}
// @lc code=end

