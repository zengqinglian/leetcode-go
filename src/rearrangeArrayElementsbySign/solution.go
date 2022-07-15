package rearrangeArrayElementsbySign

/*
Runtime: 361 ms, faster than 57.45% of Go online submissions for Rearrange Array Elements by Sign.
Memory Usage: 13.2 MB, less than 80.85% of Go online submissions for Rearrange Array Elements by Sign.
*/
func rearrangeArray(nums []int) []int {
	length := len(nums)
	numPositive := make([]int, length/2)
	numNegative := make([]int, length/2)
	indexPositive := 0
	indexNegtive := 0
	for _, n := range nums {
		if n > 0 {
			numPositive[indexPositive] = n
			indexPositive++
		} else {
			numNegative[indexNegtive] = n
			indexNegtive++
		}
	}
	indexPositive = 0
	indexNegtive = 0
	for i := range nums {
		if i%2 == 0 {
			nums[i] = numPositive[indexPositive]
			indexPositive++
		} else {
			nums[i] = numNegative[indexNegtive]
			indexNegtive++
		}
	}
	return nums
}
