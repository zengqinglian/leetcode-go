package countWordsObtainedAfterAddingaLetter

/*
Time limit over solution
failed
91 / 93 test cases passed.
*/
func wordCount1(startWords []string, targetWords []string) int {
	targetWordsArrays := make([][]int, len(targetWords))
	for i, v := range targetWords {
		wordArray := []rune(v)
		targetWordsArrays[i] = make([]int, 26)
		for _, c := range wordArray {
			targetWordsArrays[i][c-'a'] = targetWordsArrays[i][c-'a'] + 1
		}
	}
	startWordsArrays := make([][]int, len(startWords))
	for i, v := range startWords {
		wordArray := []rune(v)
		startWordsArrays[i] = make([]int, 26)
		for _, c := range wordArray {
			startWordsArrays[i][c-'a'] = startWordsArrays[i][c-'a'] + 1
		}
	}
	res := 0
	for x, t := range targetWordsArrays {
		for y, s := range startWordsArrays {
			if len(startWords[y])+1 != len(targetWords[x]) {
				continue
			}
			found := true
			for i := 0; i < 26; i++ {
				if t[i] < s[i] {
					found = false
					break
				}
			}
			if found {
				res++
				break
			}
		}
	}
	return res
}

/*
Runtime: 270 ms, faster than 28.98% of Go online submissions for Count Words Obtained After Adding a Letter.
Memory Usage: 17.3 MB, less than 44.93% of Go online submissions for Count Words Obtained After Adding a Letter.
*/
func wordCount(startWords []string, targetWords []string) int {
	startWordsSet := make(map[int]struct{})
	for _, v := range startWords {
		wordArray := []rune(v)
		total := 0
		for _, c := range wordArray {
			total |= (1 << (c - 'a'))
		}
		startWordsSet[total] = struct{}{}
	}
	targetTotal := 0
	res := 0
	for _, v := range targetWords {
		targetTotal = 0
		wordArray := []rune(v)
		for _, c := range wordArray {
			targetTotal |= (1 << (c - 'a'))
		}
		for x := range v {
			newTargetTotal := targetTotal ^ (1 << (v[x] - 'a'))
			if _, ok := startWordsSet[newTargetTotal]; ok {
				res++
				break
			}
		}
	}
	return res
}
