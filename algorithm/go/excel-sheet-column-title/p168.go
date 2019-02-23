package main

func convertToTitle(n int) string {
	arr := make([]byte, 0)
	for n > 0 {
		rest := n % 26
		if rest == 0 {
			rest = 26
		}
		arr = append(arr, byte('A'+(rest-1)))
		n = (n - rest) / 26
	}
	//reverse arr
	for i := 0; i <= (len(arr)-1)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}
