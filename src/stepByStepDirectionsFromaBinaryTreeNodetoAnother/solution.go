package stepByStepDirectionsFromaBinaryTreeNodetoAnother

import "bytes"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Path struct {
	TreeNode
	Path string
}

/*
317 / 332 test cases passed. Time over limit
*/
func getDirections(root *TreeNode, startValue int, destValue int) string {
	if startValue == destValue {
		return ""
	}

	startValuePath := getDirectionsForTarget(root, startValue)
	destValuePath := getDirectionsForTarget(root, destValue)

	samePath := true
	indexI := -1
	indexJ := -1
	for i, j := len(startValuePath)-1, len(destValuePath)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		startPathNode := startValuePath[i]
		destValuePathNode := destValuePath[j]
		indexI = i
		indexJ = j
		if startPathNode.Val != destValuePathNode.Val {
			samePath = false
			break
		}
	}
	res := ""
	if samePath {
		if len(startValuePath) < len(destValuePath) {
			//start point is parent
			for i := len(destValuePath) - len(startValuePath); i >= 0; i-- {
				res += destValuePath[i].Path
			}
		} else {
			//start point is child
			for i := 0; i < len(startValuePath)-len(destValuePath); i++ {
				res += "U"
			}
		}

	} else {
		for i := 0; i <= indexI; i++ {
			res += "U"
		}
		for i := indexJ + 1; i >= 0; i-- {
			res += destValuePath[i].Path
		}
	}

	return res
}

/*
Runtime: 539 ms, faster than 7.43% of Go online submissions for Step-By-Step Directions From a Binary Tree Node to Another.
Memory Usage: 52.2 MB, less than 9.46% of Go online submissions for Step-By-Step Directions From a Binary Tree Node to Another.
*/
func getDirections2(root *TreeNode, startValue int, destValue int) string {
	if startValue == destValue {
		return ""
	}

	startValuePath := getDirectionsForTarget(root, startValue)
	destValuePath := getDirectionsForTarget(root, destValue)

	dict := make(map[int]int)
	for i, p := range startValuePath {
		dict[p.Val] = i
	}

	var res bytes.Buffer
	for i, p := range destValuePath {
		if j, ok := dict[p.Val]; ok {
			for k := 0; k < j; k++ {
				res.WriteString("U")
			}
			for k := i; k >= 0; k-- {
				res.WriteString(destValuePath[k].Path)
			}
			break
		}
	}

	return res.String()
}

func getDirectionsForTarget(root *TreeNode, stringTarget int) []Path {
	if root.Val == stringTarget {
		pathSlice := []Path{Path{*root, ""}}
		return pathSlice
	}

	if root.Left != nil {
		subPaths := getDirectionsForTarget(root.Left, stringTarget)
		if subPaths != nil {
			return append(subPaths, Path{*root, "L"})
		}
	}

	if root.Right != nil {
		subPaths := getDirectionsForTarget(root.Right, stringTarget)
		if subPaths != nil {
			return append(subPaths, Path{*root, "R"})
		}
	}
	return nil

}

/*
Someone's quicker solution
idea is based on @votrubac's excellent post here: https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/discuss/1612105/3-Steps
Runtime: 306 ms, faster than 73.65% of Go online submissions for Step-By-Step Directions From a Binary Tree Node to Another.
Memory Usage: 33.5 MB, less than 77.03% of Go online submissions for Step-By-Step Directions From a Binary Tree Node to Another.
*/
func getDirections3(root *TreeNode, startValue int, destValue int) string {
	pathToStart := make([]byte, 0)
	find(root, startValue, &pathToStart)

	pathToDest := make([]byte, 0)
	find(root, destValue, &pathToDest)

	// remove common prefix
	i, j := 0, 0
	for i < len(pathToStart) && j < len(pathToDest) && pathToStart[i] == pathToDest[j] {
		i++
		j++
	}

	// replace pathToStart remaining bytes with U
	for k := 0; k < len(pathToStart); k++ {
		pathToStart[k] = 'U'
	}

	// concat and return the res
	res := append(pathToStart[i:], pathToDest[j:]...)

	return string(res)
}

func find(node *TreeNode, target int, res *[]byte) bool {
	if node == nil {
		return false
	}
	if node.Val == target {
		return true
	}

	// try left
	*res = append(*res, 'L')
	if find(node.Left, target, res) {
		return true
	}

	// try right
	(*res)[len(*res)-1] = 'R'
	if find(node.Right, target, res) {
		return true
	}

	*res = (*res)[:len(*res)-1]
	return false
}
