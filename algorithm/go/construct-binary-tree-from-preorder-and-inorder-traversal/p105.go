package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	return helper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func helper(pre []int, start1, end1 int, in []int, start2, end2 int) *TreeNode {
	if start1 > end1 {
		return nil
	}

	root := &TreeNode{pre[start1], nil, nil}
	if start1 == end1 {
		return root
	}

	leftLen := 0
	// rightLen := 0
	for i := start2; i <= end2; i++ {
		if in[i] == root.Val {
			leftLen = i - start2
			break
		}
	}
	// rightLen = end1 - start1 - leftLen
	root.Left = helper(pre, start1+1, start1+leftLen, in, start2, start2+leftLen-1)
	root.Right = helper(pre, start1+leftLen+1, end1, in, start2+leftLen+1, end2)
	return root

}
