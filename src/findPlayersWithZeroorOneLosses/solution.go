package findPlayersWithZeroorOneLosses

import "sort"

/*
Runtime: 423 ms, faster than 67.04% of Go online submissions for Find Players With Zero or One Losses.
Memory Usage: 26.7 MB, less than 27.27% of Go online submissions for Find Players With Zero or One Losses.
*/
func findWinners(matches [][]int) [][]int {
	loseMatchMap := make(map[int]int)
	for _, match := range matches {
		if _, ok := loseMatchMap[match[0]]; !ok {
			loseMatchMap[match[0]] = 0
		}
		if cnt, ok := loseMatchMap[match[1]]; !ok {
			loseMatchMap[match[1]] = 1
		} else {
			loseMatchMap[match[1]] = cnt + 1
		}
	}

	allwins := make([]int, 0)
	lostOne := make([]int, 0)

	for key, value := range loseMatchMap {
		if value == 0 {
			allwins = append(allwins, key)
		} else if value == 1 {
			lostOne = append(lostOne, key)
		}
	}
	sort.Ints(allwins)
	sort.Ints(lostOne)
	res := make([][]int, 2)
	res[0] = allwins
	res[1] = lostOne
	return res

}
