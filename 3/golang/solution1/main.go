package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	var flag [100]int
	n := len(s)
	result := 0
	max := 0
	for i := 0; i < n; i++ {
		char := s[i] - 31
		if max < result {
			max = result
		}
		if flag[char] != 0 {
			i = i - result + 1
			result = 0
			setZero(&flag)
			char = s[i] - 31
		}
		flag[char] = 1
		result += 1
	}
	if max < result {
		max = result
	}
	return max
}
func setZero(num *[100]int) {
	for i := 0; i < len(num); i++ {
		num[i] = 0
	}
}
