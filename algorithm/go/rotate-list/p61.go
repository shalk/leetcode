package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}

	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}
	k = k % l
	// k < l
	if k == 0 {
		return head
	}
	var ret *ListNode
	// a1 ... a2  b1 ... b2 => b1...b2 a1...a2
	a1, a2 := head, head
	for i := 0; i < (l-k)-1; i++ {
		a2 = a2.Next
	}
	ret = a2.Next

	b2 := a2.Next
	a2.Next = nil
	for b2.Next != nil {
		b2 = b2.Next
	}
	b2.Next = a1

	return ret

}
