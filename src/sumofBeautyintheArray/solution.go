package sumofBeautyintheArray

/*
Runtime: 153 ms, faster than 20.00% of Go online submissions for Sum of Beauty in the Array.
Memory Usage: 10.7 MB, less than 20.00% of Go online submissions for Sum of Beauty in the Array.
*/
func sumOfBeauties(nums []int) int {
	length := len(nums)
	suffixMin := make([]int, length)
	prefixMax := make([]int, length)
	prefixMax[0] = nums[0]
	for i := 1; i < length; i++ {
		if nums[i] >= prefixMax[i-1] {
			prefixMax[i] = nums[i]
		} else {
			prefixMax[i] = prefixMax[i-1]
		}
	}
	suffixMin[length-1] = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		if nums[i] <= suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}
	res := 0
	for i := 1; i < length-1; i++ {
		if nums[i] > prefixMax[i-1] && nums[i] < suffixMin[i+1] {
			res += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			res += 1
		}

	}
	return res
}
