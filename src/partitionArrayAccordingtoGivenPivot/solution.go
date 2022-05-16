package main

import (
	"fmt"
	"sort"
)

/*
Runtime: 365 ms, faster than 27.27% of Go online submissions for Partition Array According to Given Pivot.
Memory Usage: 12.4 MB, less than 9.09% of Go online submissions for Partition Array According to Given Pivot.
*/
func pivotArray(nums []int, pivot int) []int {

	sort.SliceStable(nums, func(i, j int) bool {
		if nums[i] < pivot && nums[j] < pivot {
			return i < j
		} else if nums[i] > pivot && nums[j] > pivot {
			return i < j
		} else {
			return nums[i] < nums[j]
		}
	})

	return nums
}

/*
Someone 's solution

Runtime: 269 ms, faster than 40.91% of Go online submissions for Partition Array According to Given Pivot.
Memory Usage: 10 MB, less than 63.64% of Go online submissions for Partition Array According to Given Pivot.
*/
func pivotArray2(nums []int, p int) []int {
	var idx int
	res := make([]int, len(nums))
	for i := range nums {
		if nums[i] < p {
			res[idx] = nums[i]
			idx++
		}
	}
	for i := range nums {
		if nums[i] == p {
			res[idx] = p
			idx++
		}
	}
	for i := range nums {
		if nums[i] > p {
			res[idx] = nums[i]
			idx++
		}
	}

	return res
}

func main() {
	nums := []int{9, 12, 5, 10, 14, 3, 10}
	fmt.Println(pivotArray(nums, 10))
}
