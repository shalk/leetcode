package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*TreeNode, 0)
	for p := root; p != nil; p = p.Left {
		stack = append(stack, p)
	}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		ret = append(ret, p.Val)
		// pop p
		stack = stack[:len(stack)-1]
		if p.Right != nil {
			for p = p.Right; p != nil; p = p.Left {
				stack = append(stack, p)
			}
		}

	}
	return ret
}
