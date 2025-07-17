/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0{
		return head
	} 
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	tail.Next = head
	step := n - k%n
	newTail := head
	for i:= 1; i < step; i++{
		newTail = newTail.Next
	}
	resultHead := newTail.Next
	newTail.Next = nil
	return resultHead
}
// @lc code=end

