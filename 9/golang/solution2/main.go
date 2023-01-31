package main

func main() {

}

/*
逆序重建了一个数、判断是否相等
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	// 余数
	rev := x % 10
	// 左边的值
	left := x / 10
	// 如果左边的值还是比右边大，继续比较
	for left > rev {
		// 余数进位
		rev = rev * 10
		// 处理下一位
		rev = rev + left%10
		// 左边减少一位
		left = left / 10
	}
	// 如果相等则是
	if left == rev {
		return true
	}
	// 如果不是则应该差了 10 倍
	if rev/10 == left {
		return true
	}
	return false
}
