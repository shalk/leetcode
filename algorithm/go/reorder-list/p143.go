package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	p1 := head
	n1 := (n + 1) / 2
	p2 := head
	// move n1-1 step
	for i := 0; i < n1-1; i++ {
		p2 = p2.Next
	}
	// disconnect p1 with p2
	tmp1 := p2.Next
	p2.Next = nil
	p2 = tmp1
	// reverse p2
	var pre *ListNode
	for p2.Next != nil {
		tmp := p2.Next
		p2.Next = pre
		pre = p2
		p2 = tmp
	}
	p2.Next = pre

	for p2 != nil {
		tmp1 := p1.Next
		tmp2 := p2.Next
		p2.Next = p1.Next
		p1.Next = p2
		p1 = tmp1
		p2 = tmp2
	}
	return
}
