package appendKIntegersWithMinimalSum

/*
Runtime: 160 ms, faster than 85.45% of Go online submissions for Append K Integers With Minimal Sum.
Memory Usage: 11.3 MB, less than 23.64% of Go online submissions for Append K Integers With Minimal Sum.
*/
func minimalKSum(nums []int, k int) int64 {
	total := (int64(1) + int64(k)) * int64(k) / int64(2)
	appeared := make(map[int]struct{})
	maxNum := k + 1
	for i := 0; i < len(nums); i++ {
		if _, ok := appeared[nums[i]]; !ok && nums[i] <= maxNum {
			total -= int64(nums[i])
			for {
				if _, ok := appeared[maxNum]; ok {
					maxNum++
				} else {
					break
				}
			}
			total += int64(maxNum)
			maxNum++
		}
		appeared[nums[i]] = struct{}{}
	}
	return total
}
