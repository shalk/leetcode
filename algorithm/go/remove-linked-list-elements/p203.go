package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	for pre, p := dummy, head; p != nil; p = p.Next {
		if p.Val == val {
			pre.Next = p.Next
			// p.Next = nil
			p = pre
		} else {
			pre = p
		}
	}
	return dummy.Next
}
