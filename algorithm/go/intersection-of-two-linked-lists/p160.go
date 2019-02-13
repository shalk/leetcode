package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	n1, tail1 := size(headA)
	n2, tail2 := size(headB)

	if tail1 != tail2 {
		return nil
	}

	for n1 != n2 {
		if n1 > n2 {
			headA = headA.Next
			n1--
		} else {
			headB = headB.Next
			n2--
		}
	}
	for headA != nil {
		if headA == headB {
			return headA
		} else {
			headA = headA.Next
			headB = headB.Next
		}
	}
	return nil

}

//return size and tail
func size(head *ListNode) (int, *ListNode) {
	n := 1
	p := head
	for p.Next != nil {
		n++
		p = p.Next
	}
	return n, p
}
