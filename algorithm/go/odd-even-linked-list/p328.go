package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	// 0/1/2 keep
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	dummyEvenHead := &ListNode{}
	p, p1 := head, dummyEvenHead
	for p.Next != nil {
		p1.Next = p.Next
		p.Next = p.Next.Next
		p1 = p1.Next
		p1.Next = nil
		if p.Next != nil {
			p = p.Next
		}
	}
	p.Next = dummyEvenHead.Next
	return dummy.Next

}
