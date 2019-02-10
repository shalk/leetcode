package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head
	for p1 != nil && p2 != nil {
		p1 = p1.Next
		if p1 == nil {
			break
		}
		p2 = p2.Next
		if p2 == nil {
			break
		}
		p2 = p2.Next
		if p2 == nil {
			break
		}
		if p1 == p2 {
			break
		}
	}
	if p1 == nil || p2 == nil {
		return nil
	}
	p1 = head
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
