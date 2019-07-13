package trees

import "math"

//https://leetcode.com/problems/maximum-depth-of-binary-tree/submissions/
//[3,9,20,null,null,15,7]
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1 + math.Max(
		float64(maxDepth1(root.Left)),
		float64(maxDepth1(root.Right)),
	)
	return int(depth)
}

func maxDepth(root *TreeNode) int {
	depth := 0.0
	if root == nil {
		return int(depth)
	}

	var stack []*TreeNode
	var el = root
	stack = push(stack, root)
	for {
		if len(stack) == 0 {
			return int(depth)
		}
		depth++

		var a = []*TreeNode{}
		for len(stack) > 0 {
			stack, el = pop(stack)
			a = append(a, el)
		}
		for _, n := range a {
			if n.Left != nil {
				stack = push(stack, n.Left)
			}
			if n.Right != nil {
				stack = push(stack, n.Right)
			}
		}
	}
}
func push(s []*TreeNode, v *TreeNode) []*TreeNode {
	return append(s, v)
}
func pop(s []*TreeNode) ([]*TreeNode, *TreeNode) {
	var el = s[len(s)-1]
	return s[:len(s)-1], el
}
