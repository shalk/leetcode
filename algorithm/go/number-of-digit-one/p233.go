package main

func countDigitOne(n int) int {
	// 1 2,3,4,5,6,8 9,0
	if n <= 0 {
		return 0
	}
	ret := 0

	prefix := n / 10
	subfix := 0
	base := 1
	for n > 0 {
		cur := n % 10
		count := 0

		if cur > 1 {
			count = (prefix + 1) * base
		} else if cur == 1 {
			count = prefix*base + subfix + 1
		} else if cur == 0 {
			count = prefix * base
		}

		ret = ret + count
		prefix /= 10
		subfix = cur*base + subfix
		base *= 10
		n = n / 10

	}
	return ret

}
