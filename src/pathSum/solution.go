package main

import "fmt"

/*
Runtime: 8 ms, faster than 40.33% of Go online submissions for Path Sum.
Memory Usage: 4.6 MB, less than 91.02% of Go online submissions for Path Sum.
Next challenges
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	res := false
	if root.Left != nil {
		res = hasPathSum(root.Left, targetSum-root.Val)
	}
	if res {
		return res
	}
	if root.Right != nil {
		res = hasPathSum(root.Right, targetSum-root.Val)
	}

	return res
}

func main() {
	fmt.Println(hasPathSum(nil, 5))
}
