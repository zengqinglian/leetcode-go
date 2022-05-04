package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 3 ms, faster than 37.63% of Go online submissions for Binary Tree Right Side View.
Memory Usage: 2.9 MB, less than 9.95% of Go online submissions for Binary Tree Right Side View.
*/
func rightSideView(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	q := make(chan TreeNode, 101)

	q <- *root

	for len(q) > 0 {
		loop := len(q)
		var node TreeNode
		for i := 0; i < loop; i++ {
			node = <-q
			if node.Left != nil {
				q <- *node.Left
			}
			if node.Right != nil {
				q <- *node.Right
			}
		}
		res = append(res, node.Val)
	}

	return res

}

func main() {
	rightSideView(nil)
}
