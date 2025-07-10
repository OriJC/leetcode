/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var s1 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	var s2 []int
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}


	dummy := &ListNode{-1, nil}
	carry := 0
	for len(s1) > 0 || len(s2) > 0 || carry > 0 {
		val := carry
		if len(s1) > 0{
			val += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) > 0{
			val += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		carry = val / 10
		val = val % 10

		node := &ListNode{Val: val, Next: dummy.Next}
		dummy.Next = node
	}
	return dummy.Next
}
// @lc code=end

