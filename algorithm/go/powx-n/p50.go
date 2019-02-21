package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	if n == 2 {
		return x * x
	}
	if n == -1 {
		return 1 / x
	}
	if x == 1.0 {
		return 1.0
	}
	if x == -1.0 {
		if n%2 == 1 {
			return -1.0
		} else {
			return 1.0
		}
	}

	if n < 0 {
		if n == math.MinInt32 {
			return myPow(1.0/x, -(n+1)) * (1.0 / x)
		}
		return myPow(1.0/x, -n)
	}
	m := n / 2
	r := myPow(x, m)
	if m*2 == n {
		return r * r
	} else {
		return r * r * x
	}
	ret := 1.0
	tmp := x
	for i := 1; i <= 31; i++ {
		if n&(1<<uint(i-1)) == (1 << uint(i-1)) {
			ret = ret * tmp
		}
		tmp = tmp * tmp
	}
	return ret
}
