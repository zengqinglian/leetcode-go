package maximizetheTopmostElementAfterKMoves

/*
Runtime: 116 ms, faster than 42.11% of Go online submissions for Maximize the Topmost Element After K Moves.
Memory Usage: 9.7 MB, less than 36.84% of Go online submissions for Maximize the Topmost Element After K Moves.
*/
func maximumTop(nums []int, k int) int {
	if k == 0 {
		return nums[0]
	}

	if len(nums) == 1 {
		if k%2 == 0 {
			return nums[0]
		} else {
			return -1
		}
	}

	res := -1
	if k <= len(nums) {
		for i := 0; i < k-1; i++ {
			if res < nums[i] {
				res = nums[i]
			}
		}

		if k < len(nums) && res < nums[k] {
			res = nums[k]
		}
		return res
	} else {
		for i := 0; i < len(nums); i++ {
			if res < nums[i] {
				res = nums[i]
			}

		}
		return res
	}

}
