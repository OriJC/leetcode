/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */

// @lc code=start
func kthSmallest(matrix [][]int, k int) int {
    pq := &IntHeap{}
	heap.Init(pq)

	for i := range matrix {
		heap.Push(pq, []int{matrix[i][0], i, 0})
	}

	res := -1
	for pq.Len() >0 && k > 0{
		cur :=  heap.Pop(pq).([]int)
		res = cur[0]
		k--

		i, j := cur[1], cur[2]

		if j + 1 < len(matrix[i]){
			heap.Push(pq, []int{matrix[i][j+1], i, j+1})
		}
	}
	return res
}

type IntHeap [][]int

func(h IntHeap) Len() int {return len(h)}

func(h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0]}

func(h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

func(h *IntHeap) Push(x interface{}){
	*h = append(*h, x.([]int))
}

func(h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}
// @lc code=end

