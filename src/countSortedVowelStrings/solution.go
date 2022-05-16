package main

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Count Sorted Vowel Strings.
Memory Usage: 1.9 MB, less than 90.10% of Go online submissions for Count Sorted Vowel Strings.
*/
func countVowelStrings(n int) int {
	values := [5]int{1, 1, 1, 1, 1}
	total := 0
	for i := 1; i <= n; i++ {
		total = 0
		for j := 4; j >= 0; j-- {
			total += values[j]
			values[j] = total
		}
	}
	return total
}
