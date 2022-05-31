package sorttheJumbledNumbers

import "sort"

/*
Runtime: 560 ms, faster than 18.18% of Go online submissions for Sort the Jumbled Numbers.
Memory Usage: 8 MB, less than 63.64% of Go online submissions for Sort the Jumbled Numbers.
*/
func sortJumbled(mapping []int, nums []int) []int {
	valueMap := make(map[int]int)
	sort.SliceStable(nums, func(i, j int) bool {
		newI := 0
		valueI := nums[i]
		if v1, ok1 := valueMap[nums[i]]; ok1 {
			newI = v1
		} else {
			power := 1
			if valueI == 0 {
				newI = mapping[0]
			} else {
				for valueI > 0 {
					newI = mapping[valueI%10]*power + newI
					valueI = valueI / 10
					power = power * 10
				}
			}
			valueMap[nums[i]] = newI
		}

		newJ := 0
		valueJ := nums[j]
		if v2, ok2 := valueMap[nums[j]]; ok2 {
			newJ = v2
		} else {
			power := 1
			if valueJ == 0 {
				newJ = mapping[0]
			} else {
				for valueJ > 0 {
					newJ = mapping[valueJ%10]*power + newJ
					valueJ = valueJ / 10
					power = power * 10
				}
			}
			valueMap[nums[j]] = newJ
		}
		return newI < newJ

	})
	return nums
}
