package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre := head
	for p := head.Next; p != nil; p = p.Next {
		if p.Val == pre.Val {
			pre.Next = p.Next
		} else {
			pre = p
		}
	}
	return head
}
