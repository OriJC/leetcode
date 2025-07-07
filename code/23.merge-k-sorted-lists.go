/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

type PQ []*ListNode

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	dummy := &ListNode{Val: -1}
	p := dummy

	pq := &PQ{}
	heap.Init(pq)

	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}
	}

	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)

		p.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}

		p = p.Next
	}
	return dummy.Next
}

// @lc code=end

