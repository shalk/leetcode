package main

func calculate(s string) int {
	numStack := make([]int, 0)
	opStack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == ' ' {

		} else if ch == '(' || ch == '+' || ch == '-' {
			opStack = append(opStack, ch)
		} else if ch == ')' {
			opStack = opStack[0 : len(opStack)-1]
			num := numStack[len(numStack)-1]
			numStack = numStack[0 : len(numStack)-1]
			if len(opStack) > 0 {
				op := opStack[len(opStack)-1]
				if op == '+' {
					// num1 = numStack[len(numStack)-1]
					numStack[len(numStack)-1] += num
					opStack = opStack[0 : len(opStack)-1]
				} else if op == '-' {
					numStack[len(numStack)-1] -= num
					opStack = opStack[0 : len(opStack)-1]
				} else {
					numStack = append(numStack, num)
				}
			} else {
				numStack = append(numStack, num)
			}
		} else {
			// ch == '0' ~ '9'
			num := int(ch - '0')
			for i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
				num = num*10 + int(s[i+1]-'0')
				i++
			}
			if len(opStack) > 0 {
				op := opStack[len(opStack)-1]
				if op == '+' {
					// num1 = numStack[len(numStack)-1]
					numStack[len(numStack)-1] += num
					opStack = opStack[0 : len(opStack)-1]
				} else if op == '-' {
					numStack[len(numStack)-1] -= num
					opStack = opStack[0 : len(opStack)-1]
				} else {
					numStack = append(numStack, num)
				}
			} else {
				numStack = append(numStack, num)
			}
		}
	}
	return numStack[0]
}
