package golang

func longestPalindrome(s string) string {
	var flag [128]int
	n := len(s)
	var max int
	var result int
	for i := 0; i < n; i++ {
		if i != 0 {
			flag[s[i]] = 0
		}
		if n-i < max {
			break
		}
		for j := i; j < n; j++ {
			if flag[s[j]] == 0 {
				flag[s[j]] = 1
			} else {
				temp1 := i
				temp2 := j
				for temp1 < temp2 {
					if s[temp1] == s[temp2] {
						temp1++
						temp2--
					} else {
						break
					}
				}
				if temp1 >= temp2 {
					temp3 := j - i + 1
					if temp3 > max {
						max = temp3
						result = i
					}
				}
			}
		}
	}
	if max == 0 {
		max++
	}
	return s[result : result+max]
}
