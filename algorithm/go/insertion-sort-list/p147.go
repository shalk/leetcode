package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	p := head.Next
	pre := head

	for p != nil {
		if pre.Val <= p.Val {
			pre = p
			p = p.Next
			continue
		}
		k := dummy
		for k.Next.Val < p.Val {
			k = k.Next
		}

		pre.Next = p.Next
		p.Next = k.Next
		k.Next = p
		p = pre.Next
	}

	return dummy.Next
}
