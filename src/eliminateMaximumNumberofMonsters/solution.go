package eliminateMaximumNumberofMonsters

import "sort"

/*
Runtime: 199 ms, faster than 50.00% of Go online submissions for Eliminate Maximum Number of Monsters.
Memory Usage: 10.5 MB, less than 25.00% of Go online submissions for Eliminate Maximum Number of Monsters.
*/
func eliminateMaximum(dist []int, speed []int) int {

	for i, j := 0, 0; i < len(dist); i, j = i+1, j+1 {
		v := dist[i] / speed[i]
		if dist[i]%speed[i] > 0 {
			v++
		}
		dist[i] = v
	}
	sort.Ints(dist)
	for i := 0; i < len(dist); i++ {
		if i >= dist[i] {
			return i
		}
	}
	return len(dist)
}
