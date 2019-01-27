package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre *ListNode
	p1 := head
	p2 := head
	// p2 move n-1
	for i := 0; i < n-1; i++ {
		p2 = p2.Next
	}
	// p1 point to 1st Node
	// p2 point to n-th Node
	// move p1 p2 together
	for p2.Next != nil {
		pre = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	// remove p1
	if pre == nil {
		return p1.Next
	} else {
		pre.Next = p1.Next
		p1.Next = nil
		return head
	}

}
