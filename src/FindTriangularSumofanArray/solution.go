package main

/*
Runtime: 72 ms, faster than 57.14% of Go online submissions for Find Triangular Sum of an Array.
Memory Usage: 5 MB, less than 56.12% of Go online submissions for Find Triangular Sum of an Array.
*/
func triangularSum(nums []int) int {
	arrayLen := len(nums)
	if arrayLen == 1 {
		return nums[0]
	}
	for i := arrayLen; i >= 2; i -= 1 {
		for j := 0; j < i-1; j += 1 {
			nums[j] = (nums[j] + nums[j+1]) % 10
		}
	}
	return nums[0]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	triangularSum(nums)
}
