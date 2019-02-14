package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}

	p1 := head
	p2 := head.Next
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	// list is even
	if p2.Next == nil {
		p2 = p1.Next
	} else {
		p2 = p1.Next.Next
	}
	mid := p1.Next
	//reverse head to end
	var pre *ListNode
	for p := head; p != mid; {
		next := p.Next
		p.Next = pre
		pre = p
		p = next
	}
	p1 = pre

	// compare
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}
