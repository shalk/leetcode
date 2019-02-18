package main

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		} else {
			return 0
		}
	}
	if dividend == 0 {
		return 0
	}
	if divisor == 2 {
		return dividend >> 1
	}
	if divisor == -2 {
		return -(dividend >> 1)
	}
	ret := 0
	flag := 1
	if divisor > 0 && dividend < 0 {
		flag = -1
	} else if divisor < 0 && dividend > 0 {
		flag = -1
	}
	if divisor < 0 {
		divisor = -divisor
	}

	if dividend == math.MinInt32 {
		dividend = dividend + divisor
		ret++
	}
	if dividend < 0 {
		dividend = -dividend
	}
	// 同时除以2
	for divisor&1 == 0 && dividend&1 == 0 {
		divisor = divisor >> 1
		dividend = dividend >> 1
	}
	for dividend >= divisor {
		var tmp uint
		a := dividend
		for (a >> 1) >= divisor {
			a = a >> 1
			tmp++
		}
		dividend = dividend - (divisor << tmp)
		ret = ret + 1<<tmp
	}
	if flag == 1 {
		return ret
	} else {
		return -ret
	}
}
