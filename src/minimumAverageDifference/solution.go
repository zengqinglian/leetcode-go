package minimumAverageDifference

import "math"

/*
Runtime: 110 ms, faster than 93.94% of Go online submissions for Minimum Average Difference.
Memory Usage: 8.3 MB, less than 91.92% of Go online submissions for Minimum Average Difference.
*/
func minimumAverageDifference(nums []int) int {
	l := len(nums)
	if l == 1 {
		return 0
	}
	for i := 1; i < l; i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	minVal := math.MaxInt
	res := -1
	for i := 0; i < l; i++ {
		div := 0
		if i != l-1 {
			div = (nums[l-1] - nums[i]) / (l - i - 1)
		}
		v := abs(nums[i]/(i+1) - div)
		if minVal > v {
			res = i
			minVal = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
