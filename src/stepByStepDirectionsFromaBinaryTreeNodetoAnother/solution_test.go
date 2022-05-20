package stepByStepDirectionsFromaBinaryTreeNodetoAnother

import "testing"

func TestFunc(t *testing.T) {
	// n5 := TreeNode{
	// 	Val: 5,
	// }

	n1 := TreeNode{
		Val: 1,
	}

	n2 := TreeNode{
		Val: 2,
	}

	// n3 := TreeNode{
	// 	Val: 3,
	// }

	// n4 := TreeNode{
	// 	Val: 4,
	// }

	// n6 := TreeNode{
	// 	Val: 6,
	// }

	// n5.Left = &n1
	// n5.Right = &n2
	// n1.Left = &n3
	// n2.Left = &n6
	// n2.Right = &n4

	n1.Left = &n2

	actual := getDirections(&n1, 2, 1)
	expected := "U"
	if actual != expected {
		t.Error("failed")
	}

}
