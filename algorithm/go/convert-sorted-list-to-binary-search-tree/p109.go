package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	if n == 0 {
		return nil
	}

	arr := make([]*ListNode, 0)
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p)
	}
	return helper(arr, 0, len(arr)-1)
}

func helper(arr []*ListNode, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return &TreeNode{arr[start].Val, nil, nil}
	}
	mid := (start + end) / 2
	root := &TreeNode{arr[mid].Val, nil, nil}
	root.Left = helper(arr, start, mid-1)
	root.Right = helper(arr, mid+1, end)
	return root
}
