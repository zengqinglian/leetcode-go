package minimumConsecutiveCardstoPickUp

import "math"

/*
Runtime: 294 ms, faster than 28.36% of Go online submissions for Minimum Consecutive Cards to Pick Up.
Memory Usage: 13.7 MB, less than 50.00% of Go online submissions for Minimum Consecutive Cards to Pick Up.

Runtime: 200 ms, faster than 77.61% of Go online submissions for Minimum Consecutive Cards to Pick Up.
Memory Usage: 14.7 MB, less than 20.90% of Go online submissions for Minimum Consecutive Cards to Pick Up.

*/
func minimumCardPickup(cards []int) int {
	valueIndexMap := make(map[int]int)
	res := math.MaxInt32
	for i, v := range cards {
		if idx, ok := valueIndexMap[v]; ok {
			dis := i - idx + 1
			if res > dis {
				res = dis
			}
		}
		valueIndexMap[v] = i
	}
	if res == math.MaxInt32 {
		return -1
	}

	return res
}
