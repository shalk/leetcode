package main

func addBinary(a string, b string) string {
	n1 := len(a)
	n2 := len(b)
	carry := 0
	buf := make([]byte, 0)
	index1 := n1 - 1
	index2 := n2 - 1
	for index1 >= 0 || index2 >= 0 {
		sum := carry
		if index1 >= 0 {
			if a[index1] == '1' {
				sum++
			}
		}
		if index2 >= 0 {
			if b[index2] == '1' {
				sum++
			}
		}
		index1--
		index2--
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		if sum == 1 {
			buf = append(buf, '1')
		} else {
			buf = append(buf, '0')
		}
	}
	if carry == 1 {
		buf = append(buf, '1')
	}
	// reverse buf
	for i := 0; i <= (len(buf)-1)/2; i++ {
		buf[i], buf[len(buf)-1-i] = buf[len(buf)-1-i], buf[i]
	}
	return string(buf)
}
