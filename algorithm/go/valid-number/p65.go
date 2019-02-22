package main

func isNumber(s string) bool {
	arr := []byte(s)
	// trim
	arr = trim(arr)
	if len(arr) == 0 {
		return false
	}
	if string(arr) == "e" {
		return false
	}
	strs := strings.Split(string(arr), "e")
	if len(strs) == 2 {
		// not allow "123e" "e3"
		return isDecimal(strs[0]) && isEPart(strs[1])
	} else if len(strs) == 1 {
		//  not allow empty
		return isDecimal(strs[0])
	} else { // 0 or more
		return false
	}

}
func isEPart(arr string) bool {
	if len(arr) == 0 {
		return false
	}
	if arr[0] == '+' || arr[0] == '-' {
		arr = arr[1:len(arr)]
	}
	if len(arr) == 0 {
		return false
	}
	return isIntPart(arr)
}

func isDecimal(arr string) bool {
	if len(arr) == 0 {
		return false
	}
	if arr[0] == '+' || arr[0] == '-' {
		arr = arr[1:len(arr)]
	}
	if len(arr) == 0 {
		return false
	}
	if string(arr) == "." {
		return false
	}
	strs := strings.Split(string(arr), ".")
	// allow “1.”  “.1"  “001.1"
	if len(strs) == 2 {
		return isIntPart(strs[0]) && isDecimalPart(strs[1])
	} else if len(strs) == 1 {
		return isIntPart(strs[0])
	} else {
		return false
	}
}

func isIntPart(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func isDecimalPart(s string) bool {
	for _, v := range s {
		if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func trim(arr []byte) []byte {
	start, end := 0, len(arr)-1
	for start < len(arr) {
		if arr[start] == ' ' {
			start++
		} else {
			break
		}
	}
	if start == len(arr) {
		return []byte{}
	}
	for end > start {
		if arr[end] == ' ' {
			end--
		} else {
			break
		}
	}
	return arr[start : end+1]
}
