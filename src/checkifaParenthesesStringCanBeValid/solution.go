package checkifaParenthesesStringCanBeValid

func canBeValid(s string, locked string) bool {
	minLeftBracket, maxLeftBracket := 0, 0
	for i, c := range s {
		if c == '(' {
			if locked[i] == '1' {
				// we have no choice, we add '('
				minLeftBracket++
				maxLeftBracket++
			} else {
				if maxLeftBracket == 0 {
					// we have to add a '('
					minLeftBracket++
				} else if minLeftBracket > 0 {
					// we can add a ')'
					minLeftBracket--
				}
				maxLeftBracket++
			}
		} else {
			if locked[i] == '1' {
				// we have no choice, we add ')'
				if minLeftBracket > 0 {
					minLeftBracket--
				}
				maxLeftBracket--
				if maxLeftBracket < 0 {
					// game over
					return false
				}
			} else {
				if maxLeftBracket == 0 {
					// we have to add a '('
					minLeftBracket++
				} else if minLeftBracket > 0 {
					// we can add a ')'
					minLeftBracket--
				}
				maxLeftBracket++
			}
		}
		// if odds elements, minLeftBracket can't be zero, at least 1
		if (i+1)%2 != 0 {
			if minLeftBracket == 0 {
				minLeftBracket = 1
			}
		}
	}
	return minLeftBracket == 0
}
