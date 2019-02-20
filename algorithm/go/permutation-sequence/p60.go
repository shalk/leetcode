package main

func getPermutation(n int, k int) string {
	candi := make([]byte, n)
	for i := 1; i <= n; i++ {
		candi[i-1] = byte('0' + i)
	}
	ret := make([]byte, 0)
	for i := n; i >= 1; i-- {
		val := jiechen(i - 1)
		index := (k - 1) / val
		k = k - val*index
		ret = append(ret, candi[index])
		// fmt.Println(string(candi[index]))
		candi = append(append([]byte{}, candi[0:index]...), candi[index+1:len(candi)]...)
		// fmt.Println(string(candi))
	}
	return string(ret)
}

func jiechen(i int) int {
	if i == 0 {
		return 1
	}
	if i == 1 {
		return 1
	}
	return i * jiechen(i-1)
}
