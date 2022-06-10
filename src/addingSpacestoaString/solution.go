package addingSpacestoaString

/*
Runtime: 141 ms, faster than 47.83% of Go online submissions for Adding Spaces to a String.
Memory Usage: 13.4 MB, less than 91.30% of Go online submissions for Adding Spaces to a String.
*/
func addSpaces(s string, spaces []int) string {
	charArray := []rune(s)
	res := make([]rune, len(charArray)+len(spaces))

	if spaces[0] == 0 {
		res[spaces[0]] = ' '
	} else {
		copy(res[0:spaces[0]], charArray[0:spaces[0]])
		res[spaces[0]] = ' '
	}

	for i := 1; i < len(spaces); i++ {
		copy(res[spaces[i-1]+i:], charArray[spaces[i-1]:spaces[i]])
		res[spaces[i]+i] = ' '
	}
	copy(res[spaces[len(spaces)-1]+len(spaces):], charArray[spaces[len(spaces)-1]:])
	return string(res)
}
