package findPalindromeWithFixedLength

import "math"

/*
Runtime: 176 ms, faster than 50.00% of Go online submissions for Find Palindrome With Fixed Length.
Memory Usage: 10.2 MB, less than 13.64% of Go online submissions for Find Palindrome With Fixed Length.
*/
func kthPalindrome(queries []int, intLength int) []int64 {

	midIndex := intLength / 2
	if intLength%2 == 0 {
		midIndex--
	}
	total := 9 * int(math.Pow10(midIndex))

	res := make([]int64, len(queries))
	for i, query := range queries {
		if query > total {
			res[i] = -1
		} else {
			res[i] = calculate(query-1, midIndex, intLength)
		}
	}
	return res

}

func calculate(query int, digitCount int, intLength int) int64 {
	number := make([]int, intLength)
	i := 0
	length := 0
	for length+2 < intLength && query > 0 {
		v := query % 10
		if intLength%2 == 1 {
			number[digitCount-i] = v
			number[digitCount+i] = v
			if i == 0 {
				length++
			} else {
				length += 2
			}
		} else {
			number[digitCount-i] = v
			number[digitCount+1+i] = v
			length += 2
		}
		i++
		query /= 10
	}

	number[0] = query + 1
	number[len(number)-1] = query + 1

	res := int64(0)
	for _, v := range number {
		res = res*int64(10) + int64(v)
	}
	return res
}
