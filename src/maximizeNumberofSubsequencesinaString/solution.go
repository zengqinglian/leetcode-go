package maximizeNumberofSubsequencesinaString

/*
Runtime: 16 ms, faster than 64.71% of Go online submissions for Maximize Number of Subsequences in a String.
Memory Usage: 6.6 MB, less than 58.82% of Go online submissions for Maximize Number of Subsequences in a String.
*/
func maximumSubsequenceCount(text string, pattern string) int64 {
	textRune := []rune(text)
	patternRune := []rune(pattern)
	cntR0 := int64(0)
	cntR1 := int64(0)
	total := int64(0)
	for i := len(textRune) - 1; i >= 0; i-- {
		if patternRune[1] == patternRune[0] {
			if textRune[i] == patternRune[1] {
				total += cntR0
				cntR0++
			}
		} else {
			if textRune[i] == patternRune[1] {
				cntR1++
			}
			if textRune[i] == patternRune[0] {
				total += cntR1
				cntR0++
			}
		}
	}
	if cntR0 >= cntR1 {
		total += cntR0
	} else {
		total += cntR1
	}

	return total
}
