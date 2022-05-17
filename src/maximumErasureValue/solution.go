package maximumErasureValue

/*
Runtime: 151 ms, faster than 82.61% of Go online submissions for Maximum Erasure Value.
Memory Usage: 9.2 MB, less than 60.87% of Go online submissions for Maximum Erasure Value.
*/
func maximumUniqueSubarray(nums []int) int {
	array := [10001]int{}
	j := 0
	total := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		if array[nums[i]] == 0 {
			total += nums[i]
			array[nums[i]] = 1
		} else {
			for nums[j] != nums[i] {
				total -= nums[j]
				array[nums[j]] = 0
				j++
			}
			j++
		}
		if total > max {
			max = total
		}
	}
	return max
}
