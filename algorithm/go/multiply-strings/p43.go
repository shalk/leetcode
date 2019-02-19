package main

func multiply(num1 string, num2 string) string {
	if num1 == "0" {
		return "0"
	}
	if num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}
	sum := make([]int, len(num1)+len(num2))
	for index1 := len(num1) - 1; index1 >= 0; index1-- {
		for index2 := len(num2) - 1; index2 >= 0; index2-- {
			val := int(num1[index1]-'0') * int(num2[index2]-'0')
			sum[len(num1)-index1-1+len(num2)-index2-1] += val
		}
	}
	carry := 0
	for i := 0; i < len(sum); i++ {
		sum[i] += carry
		carry = 0
		carry = sum[i] / 10
		sum[i] = sum[i] - 10*carry
	}
	// reverse sum
	for i := 0; i <= (len(sum)-1)/2; i++ {
		sum[i], sum[len(sum)-1-i] = sum[len(sum)-1-i], sum[i]
	}
	ret := make([]byte, 0)
	// skip first 0
	start := 0
	if sum[0] == 0 {
		start = 1
	}
	for ; start < len(sum); start++ {
		ret = append(ret, byte(sum[start]+'0'))
	}
	return string(ret)
}
