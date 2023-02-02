package solution1

import "math"

func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	temp1 := x

	if x < 0 {
		temp1 = -x
		if temp1 <= 0 {
			return 0
		}
	}
	// 小于10
	if temp1 < 10 {
		return x
	}
	result := 0
	for temp1%10 != 0 || temp1/10 != 0 {
		result = result*10 + temp1%10
		temp1 /= 10
	}
	if result > math.MaxInt32 {
		return 0
	}
	if x < 0 {
		if -result < math.MinInt32 {
			return 0
		}
		return -result
	}
	return result
}
