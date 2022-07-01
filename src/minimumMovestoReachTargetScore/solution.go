package minimumMovestoReachTargetScore
/*
Runtime: 2 ms, faster than 48.65% of Go online submissions for Minimum Moves to Reach Target Score.
Memory Usage: 1.9 MB, less than 40.54% of Go online submissions for Minimum Moves to Reach Target Score.
*/
func minMoves(target int, maxDoubles int) int {
	res := 0
	for target > 1 && maxDoubles > 0 {
		if target%2 == 1 {
			res++
			target--
		} else {
			res++
			target /= 2
			maxDoubles--
		}
	}
	return res + target - 1
}
