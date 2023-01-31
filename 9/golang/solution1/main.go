package main

func main() {

}

func isPalindrome(x int) bool {
	// 小于0
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var temp []int
	// 如果不是一位数
	temp2 := x % 10
	for x != 0 {
		temp = append(temp, temp2)
		x = x / 10
		temp2 = x % 10
	}
	n := len(temp)
	for i := 0; i < n/2; i++ {
		if temp[i] != temp[n-1-i] {
			return false
		}
	}
	return true
}
