package countCollisionsonaRoad

/*
Runtime: 27 ms, faster than 56.25% of Go online submissions for Count Collisions on a Road.
Memory Usage: 7 MB, less than 43.75% of Go online submissions for Count Collisions on a Road.
*/
func countCollisions(directions string) int {
	strChars := []rune(directions)
	deque := make([]rune, 0, len(strChars))
	res := 0
	for _, chr := range strChars {
		switch chr {
		case 'R':
			if len(deque) == 0 {
				deque = append(deque, chr)
			} else {
				if deque[len(deque)-1] == 'S' {
					deque = deque[:0]
					deque = append(deque, chr)
				} else if deque[len(deque)-1] == 'R' {
					deque = append(deque, chr)
				}
			}
		case 'L':
			if len(deque) > 0 {
				if deque[len(deque)-1] == 'R' {
					res += 2
					res += (len(deque) - 1)
					deque = deque[:0]
					deque = append(deque, 'S')
				} else if deque[len(deque)-1] == 'S' {
					res++
				} else {
					panic("unexpected status in deque!")
				}
			}
		case 'S':
			if len(deque) == 0 {
				deque = append(deque, chr)
			} else {
				if deque[len(deque)-1] == 'R' {
					res += len(deque)
					deque = deque[:0]
					deque = append(deque, chr)
				}
			}
		}

	}
	return res
}
