package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre, p1 := dummy, head
	for p1 != nil && p1.Next != nil {
		p2 := p1.Next
		pre.Next = p2
		p1.Next = p2.Next
		p2.Next = p1
		//
		pre = p1
		p1 = pre.Next
	}
	return dummy.Next
}
