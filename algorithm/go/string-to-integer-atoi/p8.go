package main

//" 0000000000012345678"
//"3.14"
//"    0000000000000   "
//"  -0012a42"
//"words and 987"

func myAtoi(str string) int {
	first, last := -1, -1
	// find word trimed
	for i, c := range str {
		if first == -1 {
			if c != ' ' {
				first = i
				last = first
			}
		} else {
			if c == ' ' {
				break
			} else {
				last = i
			}
		}
	}

	if first == -1 {
		return 0
	}

	word := str[first : last+1]
	// check word is +/- numbers
	sign := 1
	if word[0] == '-' {
		word = word[1:len(word)]
		sign = -1
	} else if word[0] == '+' {
		word = word[1:len(word)]
		sign = 1
	}
	if len(word) == 0 {
		return 0
	}

	//if word[0] < '0' || word[0] > '9' {
	//  return 0
	//}
	// point count
	point := 0
	last = -2
	// check is number, or number.number
	for i, c := range word {
		if c < '0' || c > '9' {
			if c == '.' && point == 0 {
				point++
				last = i - 1
				break
			} else {
				last = i - 1
				break
			}
		} else {
		}
	}
	// fmt.Printf("last:%d\n", last)
	// have 3.14
	// if point == 1 {
	if last != -2 {
		word = word[0 : last+1]
		if len(word) == 0 {
			return 0
		}
	}
	// fmt.Printf("word:%s\n", word)
	// }
	// omit start zero
	start := -1
	for i, c := range word {
		if c != '0' {
			start = i
			break
		}
	}
	if start == -1 {
		return 0
	}
	word = word[start:len(word)]
	if len(word) == 0 {
		return 0
	}

	// word is too large
	if len(word) > 10 {
		if sign == 1 {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}

	}

	//convert string to number
	ret := 0
	for _, c := range word {
		ret = ret*10 + int(c-'0')
	}
	// check range
	if sign == 1 {
		if ret > math.MaxInt32 {
			return math.MaxInt32
		}
		return ret
	} else {
		if -ret < math.MinInt32 {
			return math.MinInt32
		}
		return -ret
	}
}
