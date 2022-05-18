package maximalNetworkRank

/*
Runtime: 59 ms, faster than 65.00% of Go online submissions for Maximal Network Rank.
Memory Usage: 7.7 MB, less than 45.00% of Go online submissions for Maximal Network Rank.
Next challenges:
*/
func maximalNetworkRank(n int, roads [][]int) int {
	array := make([][]int, n)
	for _, road := range roads {
		if array[road[0]] == nil {
			array[road[0]] = make([]int, 0, n)
		}
		array[road[0]] = append(array[road[0]], road[1])

		if array[road[1]] == nil {
			array[road[1]] = make([]int, 0, n)
		}
		array[road[1]] = append(array[road[1]], road[0])
	}

	max := -1
	next := -1
	for _, arr := range array {
		l := len(arr)
		if l >= max {
			next = max
			max = l
		} else if l > next {
			next = l
		}
	}
	maxIndexes := make([]int, 0, n)
	nextIndexes := make([]int, 0, n)
	for idx, arr := range array {
		l := len(arr)
		if l == max {
			maxIndexes = append(maxIndexes, idx)
		}
		if l == next {
			nextIndexes = append(nextIndexes, idx)
		}
	}

	for _, x := range maxIndexes {
		for _, y := range nextIndexes {
			if x != y {
				if calculate(x, y, array) == 0 {
					return max + next
				}
			}
		}
	}
	return max + next - 1
}

func calculate(x int, y int, array [][]int) int {
	for _, v := range array[x] {
		if v == y {
			return 1
		}
	}
	return 0
}
