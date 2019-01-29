package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	if k == 1 {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	p := dummy.Next
	for p != nil {
		first := p
		last := p
		i := 0
		for ; i < k-1 && last.Next != nil; i++ {
			last = last.Next
		}
		if i != k-1 {
			return dummy.Next
		}
		// disconnect last.Next
		next := last.Next
		last.Next = nil
		h := reverse(first)
		// reconnect the list
		pre.Next = h
		first.Next = next
		// next loop
		pre = first
		p = pre.Next
	}

	return dummy.Next
}

func reverse(h *ListNode) *ListNode {
	if h == nil {
		return h
	}
	var pre *ListNode
	for p := h; p != nil; {
		next := p.Next
		p.Next = pre
		pre = p
		if next == nil {
			return p
		}
		p = next
	}
	return nil
}
