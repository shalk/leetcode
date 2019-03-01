package main

func nthUglyNumber(n int) int {
	m := []int{1}
	l1 := 0
	l2 := 0
	l3 := 0
	for n > len(m) {
		var min int
		// fmt.Printf("%d %d %d \n", m[l1]*2, m[l2]*3, m[l3]*5)
		if m[l1]*2 <= m[l2]*3 && m[l1]*2 <= m[l3]*5 {
			min = m[l1] * 2
			l1++
		} else if m[l2]*3 <= m[l1]*2 && m[l2]*3 <= m[l3]*5 {
			min = m[l2] * 3
			l2++
		} else {
			min = m[l3] * 5
			l3++
		}
		if min > m[len(m)-1] {
			m = append(m, min)
		}

	}
	return m[n-1]
}
