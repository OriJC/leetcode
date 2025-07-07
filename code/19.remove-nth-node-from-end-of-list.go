/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head

	x := findFromEnd(dummy, n+1)

	x.Next = x.Next.Next

	return dummy.Next
}

func findFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	p2 := head

	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	return p2
}

// @lc code=end

