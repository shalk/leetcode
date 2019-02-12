package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := size(head)
	return helper(head, n)

}

func helper(p *ListNode, s int) *ListNode {
	if s <= 1 {
		return p
	}

	n1 := s / 2
	p2 := p
	for i := 0; i < n1-1; i++ {
		p2 = p2.Next
	}

	tmp := p2.Next
	p2.Next = nil
	p2 = tmp

	n2 := s - n1

	p = helper(p, n1)
	p2 = helper(p2, n2)

	p = merge(p, p2)
	return p
}

func size(head *ListNode) int {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	return n
}

func merge(p1, p2 *ListNode) *ListNode {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}
	var h *ListNode
	if p1.Val > p2.Val {
		h = p2
		p2 = p2.Next
	} else {
		h = p1
		p1 = p1.Next
	}
	p := h
	// return p1
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p = p.Next
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p.Next
			p2 = p2.Next
		}
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return h
}
