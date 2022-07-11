package countWordsObtainedAfterAddingaLetter

import "testing"

func TestFunc(t *testing.T) {
	start := []string{"ab", "a"}
	end := []string{"abc", "abcd"}
	wordCount(start, end)
}
