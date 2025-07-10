/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    pq := &PQ{}
	heap.Init(pq)

	for i := 0; i < len(nums1); i++ {
		heap.Push(pq, []int {nums1[i], nums2[0], 0})
	}

	res := [][]int {}

	for pq.Len() > 0 && k > 0{
		cur := heap.Pop(pq).([]int)
		k--

		nextIndex := cur[2] + 1

		if nextIndex < len(nums2){
			heap.Push(pq, []int{cur[0], nums2[nextIndex], nextIndex})
		}

		res = append(res, []int{cur[0], cur[1]})
	}
	return res
}

type PQ [][]int

func(h PQ) Len() int { return len(h)}

func(h PQ) Less(i, j int) bool { return (h[i][0]+ h[i][1]) < (h[j][0]+ h[j][1])}

func(h PQ) Swap(i, j int) { h[i], h[j] = h[j], h[i]}

func(h *PQ) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func(h *PQ) Pop() interface{}{
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0:n-1]
	return item
}
// @lc code=end

