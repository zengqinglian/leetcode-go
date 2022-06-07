package createBinaryTreeFromDescriptions

import "testing"

func TestFunc(t *testing.T) {
	//descriptions := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
	//descriptions := [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}}
	descriptions := [][]int{{85, 82, 1}, {74, 85, 1}, {39, 70, 0}, {82, 38, 1}, {74, 39, 0}, {39, 13, 1}}
	createBinaryTree(descriptions)
}
