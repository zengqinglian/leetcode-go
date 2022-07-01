package minimumOperationstoHalveArraySum

/*
Runtime: 299 ms, faster than 37.50% of Go online submissions for Minimum Operations to Halve Array Sum.
Memory Usage: 10.6 MB, less than 37.50% of Go online submissions for Minimum Operations to Halve Array Sum.
*/
import "container/heap"

type DoubleHeap []float64
type any = interface{}

func (h DoubleHeap) Len() int           { return len(h) }
func (h DoubleHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h DoubleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DoubleHeap) Push(x any) {
	*h = append(*h, x.(float64))
}

func (h *DoubleHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func halveArray(nums []int) int {
	var sum float64 = 0.0
	q := make(DoubleHeap, 0)
	for _, n := range nums {
		sum += float64(n)
		heap.Push(&q, float64(n))
	}
	currentSum := float64(sum)
	cnt := 0
	for currentSum*2 > sum {
		value := heap.Pop(&q)
		currentSum -= value.(float64)
		halfValue := value.(float64) / 2
		if len(q) > 0 {
			for halfValue >= q[0] && (currentSum+halfValue)*2 > sum {
				halfValue = halfValue / 2
				cnt++
			}
		}
		heap.Push(&q, halfValue)
		currentSum += halfValue
		cnt++
	}
	return cnt
}
