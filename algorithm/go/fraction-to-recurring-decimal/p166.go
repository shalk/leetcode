package main

import "fmt"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	arr := make([]byte, 0)
	if numerator < 0 && denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	} else if numerator < 0 {
		numerator = -numerator
		arr = append(arr, '-')
	} else if denominator < 0 {
		denominator = -denominator
		arr = append(arr, '-')
	} else {

	}
	if numerator >= denominator {
		val := numerator / denominator
		numerator = numerator - denominator*val
		bytes := []byte(fmt.Sprintf("%d", val))
		arr = append(arr, bytes...)
	} else {
		arr = append(arr, '0')
	}
	if numerator == 0 {
		return string(arr)
	} else {
		arr = append(arr, '.')
	}
	n := numerator
	d := denominator
	//保存余数
	m := make(map[int]int)
	m[n] = len(arr)
	for n != 0 {
		n = n * 10
		val := n / d  // 商
		n = n - d*val // 余数
		arr = append(arr, byte(val+'0'))
		if v, ok := m[n]; ok {
			arr = append(arr[:v], append([]byte{'('}, arr[v:]...)...)
			arr = append(arr, ')')
			return string(arr)
		} else {
			m[n] = len(arr)
		}
	}
	return string(arr)
}
