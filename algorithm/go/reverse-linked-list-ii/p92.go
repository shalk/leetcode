package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	for k := 0; k < m-1; k++ {
		pre = pre.Next
	}
	var pre1 *ListNode
	last := pre.Next
	p := pre.Next
	for k := 0; k < n-m+1; k++ {
		tmp := p.Next
		p.Next = pre1
		pre1 = p
		p = tmp
	}
	// pre1 record first, p point to next part
	pre.Next = pre1
	last.Next = p
	return dummy.Next
}
