package solution1

import "math"

func myAtoi(s string) int {
	flag := true
	start := false
	result := 0
	for _, char := range s {
		if char == '-' && !start {
			flag = false
			start = true
		} else if char == '+' && !start {
			flag = true
			start = true
		} else if char >= '0' && char <= '9' {
			start = true
			temp := char - '0'
			result = result*10 + int(temp)
			if result > math.MaxInt32 {
				if flag {
					return math.MaxInt32
				}
				return math.MinInt32
			}
		} else if char == ' ' && !start {
			continue
		} else if start {
			break
		} else {
			break
		}
	}
	if flag {
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
	} else {
		result = -result
		if result < math.MinInt32 {
			return math.MinInt32
		}
	}
	return result
}
