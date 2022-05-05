package main

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2 MB, less than 43.29% of Go online submissions for Valid Parentheses.
*/
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0, 10000)
	for _, c := range s {
		if c == '[' || c == '(' || c == '{' {
			stack = append(stack, c)
		} else if c == ']' {
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if c == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if c == '}' {
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	isValid("(]")
}
