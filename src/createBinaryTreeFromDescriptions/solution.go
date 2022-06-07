package createBinaryTreeFromDescriptions

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 562 ms, faster than 25.00% of Go online submissions for Create Binary Tree From Descriptions.
Memory Usage: 9.7 MB, less than 6.82% of Go online submissions for Create Binary Tree From Descriptions.
Next challenges:
*/
func createBinaryTree(descriptions [][]int) *TreeNode {
	childrenMap := make(map[int]bool)
	relationMap := make(map[int]*TreeNode)

	for _, description := range descriptions {
		var child *TreeNode
		if c, ok := relationMap[description[1]]; ok {
			child = c
		} else {
			child = &TreeNode{Val: description[1]}
			relationMap[description[1]] = child
		}

		if parent, ok := relationMap[description[0]]; !ok {
			p := TreeNode{Val: description[0]}
			if description[2] == 1 {
				p.Left = child
			} else {
				p.Right = child
			}
			relationMap[description[0]] = &p
		} else {
			if description[2] == 1 {
				parent.Left = child
			} else {
				parent.Right = child
			}
		}
		childrenMap[description[1]] = true
	}

	for key := range relationMap {
		if !childrenMap[key] {
			return relationMap[key]
		}
	}

	return nil

}

