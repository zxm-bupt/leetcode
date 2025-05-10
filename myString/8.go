package myString

import "math"

func MyAtoi(s string) int {
	var result int64 = 0
	symbol := 0
	state := 1
	flag := 0
	for i := 0; i < len(s); i++ {
		switch state {
		case 1:
			if s[i] == '-' {
				symbol = 1
				state = 2
			} else if s[i] == '+' || s[i] == '0' {
				symbol = 0
				state = 2
			} else if s[i] <= '9' && s[i] >= '1' {
				result = result*10 + int64(s[i]-'0')
				state = 3
			} else if s[i] == ' ' {
				continue
			} else {
				flag = 1
			}
		case 2:
			if s[i] <= '9' && s[i] >= '1' {
				result = result*10 + int64(s[i]-'0')
				state = 3
			} else if s[i] == '0' {
				continue
			} else {
				flag = 1
			}
		case 3:
			if s[i] <= '9' && s[i] >= '0' {
				result = result*10 + int64(s[i]-'0')
				if result > math.MaxInt32 {
					if symbol == 1 {
						return math.MinInt32
					}
					return math.MaxInt32
				}
			} else {
				flag = 1
			}
		}
		if flag == 1 {
			break
		}
	}
	if symbol == 1 {
		result = 0 - result
	}

	return int(result)
}
