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
func buildTree(inorder []int, postorder []int) *TreeNode {
	return helper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func helper(in []int, start1, end1 int, post []int, start2, end2 int) *TreeNode {
	if start1 > end1 {
		return nil
	}
	root := &TreeNode{post[end2], nil, nil}
	if start1 == end1 {
		return root
	}
	leftLen := 0
	for i := start1; i <= end1; i++ {
		if in[i] == root.Val {
			leftLen = i - start1
			break
		}
	}
	root.Left = helper(in, start1, start1+leftLen-1, post, start2, start2+leftLen-1)
	root.Right = helper(in, start1+leftLen+1, end1, post, start2+leftLen, end2-1)
	return root
}
