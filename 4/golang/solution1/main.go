package solution1

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var temp []int
	n1 := len(nums1)
	n2 := len(nums2)

	// 判断是奇数还是偶数
	flag := (n1 + n2) % 2
	num := (n1 + n2) / 2

	i, j := 0, 0
	for i+j <= num {
		if n1 == 0 {
			temp = nums2
			break
		} else if n2 == 0 {
			temp = nums1
			break
		}
		if i < n1 && j < n2 && nums1[i] <= nums2[j] {
			temp = append(temp, nums1[i])
			i++
			continue
		}
		if i < n1 && j < n2 && nums2[j] < nums1[i] {
			temp = append(temp, nums2[j])
			j++
			continue
		}
		if i == n1 {
			for k := j; k < n2; k++ {
				temp = append(temp, nums2[k])
				j++
			}
			continue
		}
		if j == n2 {
			for k := i; k < n1; k++ {
				temp = append(temp, nums1[k])
				i++
			}
			continue
		}
	}
	if flag == 1 {
		return float64(temp[num])
	}
	return (float64(temp[num-1]) + float64(temp[num])) / 2
}

/*
*
错误太多、无法修改完善
*/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)

	// 判断是奇数还是偶数
	flag := (n1 + n2) % 2
	num := (n1 + n2) / 2
	// 空数组处理
	if num == 0 {
		if n2 == 0 {
			return float64(nums1[0])
		}
		return float64(nums2[0])
	}
	if n1 == 0 || n2 == 0 {
		if flag == 1 {
			if n1 == 0 {
				return float64(nums2[num])
			}
			return float64(nums1[num])
		}
		if n1 == 0 {
			return (float64(nums2[num-1]) + float64(nums2[num])) / 2
		}
		return (float64(nums1[num-1]) + float64(nums1[num])) / 2
	}

	// 奇数与偶数应该分开处理
	// 奇数
	if flag == 1 {
		// 遍历数组
		i, j := 0, 0
		for i+j < num {
			if nums1[i] <= nums2[j] {
				i++
			} else {
				j++
			}
		}
		if i == n1 {
			i--
		}
		if j == n2 {
			j--
		}
		if nums1[i] < nums2[j] {
			return float64(nums1[i])
		}
		return float64(nums2[j])
	}
	// 偶数（需要进行中间取值）
	// 遍历数组
	i, j := 0, 0
	for i+j < num-1 {
		if nums1[i] <= nums2[j] {
			i++
		} else {
			j++
		}
	}
	if i == n1 && i != 0 {
		i--
	}
	if j == n2 && j != 0 {
		j--
	}
	return (float64(nums1[i]) + float64(nums2[j])) / 2.0
}

/*
*
空间复杂度为 O(1) 的算法，在所有的值中间找到需要两个数值一左一右
*/
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	p1, p2 := 0, 0
	c1, c2 := len(nums1), len(nums2)
	left, right := -1, -1
	length := c1 + c2
	for i := 0; i < length/2+1; i++ {
		// 从左向右查找
		left = right
		// 进入这个分支的情况：
		// 1. 数组一还未到尽头 且 (数组二已经到尽头 或 数组一当前的值比数组二当前的值小）
		if p1 < c1 && (p2 >= c2 || nums1[p1] < nums2[p2]) {
			right = nums1[p1]
			p1++
			// 进入这个分支的情况：
			// 1. 数组二未到尽头，且当前值比数组一的值大
		} else {
			right = nums2[p2]
			p2++
		}
	}
	// 偶数、求算术平均值
	if length%2 == 0 {
		return float64(left+right) / 2.0
	}
	// 奇数，现在的右值
	return float64(right)
}
