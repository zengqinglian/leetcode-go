package numberofwaystosplitarray

/*
Runtime: 168 ms, faster than 82.35% of Go online submissions for Number of Ways to Split Array.
Memory Usage: 9.6 MB, less than 82.35% of Go online submissions for Number of Ways to Split Array.
*/
func waysToSplitArray(nums []int) int {
	total := 0
	for i, n := range nums {
		total += n
		nums[i] = total
	}
	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= total-nums[i] {
			cnt++
		}
	}
	return cnt
}
