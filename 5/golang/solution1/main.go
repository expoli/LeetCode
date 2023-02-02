package solution1

func longestPalindrome1(s string) string {
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

/*
*
算法思想：选一个数、然后依据这个数同时想前与向后进行拓展
如果向前的与向后的数值一致的话，那么这个子串就是回文子串
如果后面未能发现新的子串，那么向右更新子串判断的区间
*/
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	l, r := 0, 0
	max := 0
	res := ""
	for l < n && r < n {
		if n-r < max/2 {
			break
		}
		i, j := l, r
		// 由中间向两边进行扩散比对
		for i >= 0 && j < n && s[i] == s[j] {
			i--
			j++
		}
		// 计算子串的长度
		strLen := j - i - 1
		// 更新最长子串的长度
		if strLen > max {
			max = strLen
			res = s[i+1 : j]
		}
		if l == r {
			r++
		} else {
			l++
		}
	}
	return res
}
