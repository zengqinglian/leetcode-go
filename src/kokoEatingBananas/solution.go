package main

/*
Runtime
3 ms
Beats
71.89%
Analyze Complexity
Memory
8.10 MB
Beats
27.49%
*/
func minEatingSpeed(piles []int, h int) int {
	minVal := 1
	maxVal := 1
	for _, pile := range piles {
		maxVal = max(maxVal, pile)
		minVal = min(minVal, pile)
	}
	loop := h / len(piles)

	lower := minVal / loop
	upper := maxVal / loop
	if maxVal%loop > 0 {
		upper++
	}
	if lower == 0 {
		lower++
	}
	for lower+1 < upper {
		mid := (lower + upper) / 2
		if helper(piles, mid) <= h {
			upper = mid
		} else {
			lower = mid
		}
	}
	checkAgain := helper(piles, lower)
	if checkAgain <= h {
		return lower
	}
	return upper

}

func helper(piles []int, val int) int {
	total := 0
	for _, pile := range piles {
		total += pile / val
		if pile%val > 0 {
			total++
		}
	}
	return total
}
