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
	dummy := &ListNode{}
	pre := dummy
	for p := head; p != nil; {
		if (p.Next != nil && p.Next.Val != p.Val) || p.Next == nil {
			pre.Next = p
			pre = p
			p = p.Next
			pre.Next = nil
		} else { // p.Next != nil && p.Next.Val == v
			v := p.Val
			p = p.Next
			for p != nil && p.Val == v {
				p = p.Next
			}
		}
	}
	return dummy.Next
}
