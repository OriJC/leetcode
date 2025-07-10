/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    dummyUni := &ListNode{Val: 101}
	dummyDup := &ListNode{Val: 101}

	pUni := dummyUni
	pDup := dummyDup
	p := head
	for p != nil{
		if p.Val == pDup.Val || (p.Next != nil && p.Val == p.Next.Val){
			pDup.Next = p
			pDup = pDup.Next
		} else {
			pUni.Next = p
			pUni = pUni.Next
		}

		p = p.Next
		pUni.Next = nil
		pDup.Next = nil
	}

	return dummyUni.Next
}
// @lc code=end

