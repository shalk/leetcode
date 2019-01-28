package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	// min Stack
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	stack := make([]*ListNode, 0, len(lists))
	n := 0
	// put
	for _, h := range lists {
		if h != nil {
			stack = append(stack, h)
			for i := len(stack) - 1; i > 0; {
				parent := (i - 1) / 2
				if stack[i].Val < stack[parent].Val {
					stack[i], stack[parent] = stack[parent], stack[i]
					i = parent
				} else {
					break
				}
			}
			n++
		}
	}
	dummy := &ListNode{}
	p := dummy
	for n > 0 {
		// remove min from stack;
		p.Next = stack[0]
		p = p.Next
		stack[0] = stack[0].Next
		p.Next = nil
		// balance min stack
		if stack[0] == nil {
			n--
		}
		if n == 0 {
			break
		}
		if stack[0] == nil {
			stack[len(stack)-1], stack[0] = stack[0], stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}

		for i := 0; i < len(stack); {
			// left is (i+1)*2-1 right is (i+1)*2
			left := (i+1)*2 - 1
			right := (i + 1) * 2
			if left >= len(stack) { // do not left and right
				break
			} else if left == len(stack)-1 { // only have left
				if stack[left].Val < stack[i].Val {
					stack[i], stack[left] = stack[left], stack[i]
					i = left
				} else {
					break
				}
			} else { // have left and right
				if stack[left].Val <= stack[right].Val && stack[left].Val < stack[i].Val {
					stack[i], stack[left] = stack[left], stack[i]
					i = left
				} else if stack[right].Val <= stack[left].Val && stack[right].Val < stack[i].Val {
					stack[i], stack[right] = stack[right], stack[i]
					i = right
				} else { // i is min
					break
				}

			}
		}
		// fmt.Printf("STACK:")
		// for _, v := range stack {
		//     fmt.Printf("%d,", v.Val)
		// }
		// fmt.Println()
	}
	return dummy.Next
}
