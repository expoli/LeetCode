package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	var flag [128]int
	n := len(s)
	result := 0
	rk := -1
	for i := 0; i < n; i++ {
		if i != 0 {
			flag[s[i-1]] = 0
		}
		// 打标记
		for rk+1 < n && flag[s[rk+1]] == 0 {
			flag[s[rk+1]] = 1
			rk++
		}
		result = max(result, rk-i+1)
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
