package maximumProductAfterKIncrements

/*
Runtime: 602 ms, faster than 13.72% of Go online submissions for Maximum Product After K Increments.
Memory Usage: 10.6 MB, less than 27.45% of Go online submissions for Maximum Product After K Increments.
*/
import "container/heap"

type IntHeap []int
type any = interface{}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0] + k
	}
	q := make(IntHeap, 0)
	zeroCnt := 0
	for _, v := range nums {
		heap.Push(&q, v)
		if v == 0 {
			zeroCnt++
		}
	}

	if k < zeroCnt {
		return 0
	}

	for k > 0 {
		v := heap.Pop(&q).(int)
		next := q[0]
		if v == next {
			k--
			heap.Push(&q, v+1)
		} else if next-v >= k {
			heap.Push(&q, v+k)
			k = 0
		} else {
			k -= (next + 1 - v)
			heap.Push(&q, next+1)
		}
	}
	res := int64(1)
	for len(q) > 0 {
		v := heap.Pop(&q).(int)
		res *= int64(v)
		res %= 1_000_000_007
	}
	return int(res)

}
