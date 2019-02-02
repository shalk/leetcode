package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1 := dummy1
	p2 := dummy2
	for p := head; p != nil; {
		k := p.Next
		if p.Val < x {
			p1.Next = p
			p1 = p
		} else {
			p2.Next = p
			p2 = p
		}
		p.Next = nil
		p = k
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
