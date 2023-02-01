package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	num1 := []int{1, 3}
	num2 := []int{2}

	num3 := []int{1, 2}
	num4 := []int{3, 4}

	assertions := assert.New(t)

	assertions.Equal(float64(2), findMedianSortedArrays(num1, num2))
	assertions.Equal(2.5, findMedianSortedArrays(num3, num4))

	num3 = []int{2}
	num4 = []int{}
	assertions.Equal(float64(2), findMedianSortedArrays(num3, num4))

	num3 = []int{0, 0, 0, 0, 0}
	num4 = []int{-1, 0, 0, 0, 0, 0, 1}
	assertions.Equal(float64(0), findMedianSortedArrays(num3, num4))

	num3 = []int{1, 3}
	num4 = []int{2, 7}
	assertions.Equal(float64(2.5), findMedianSortedArrays(num3, num4))

	num3 = []int{}
	num4 = []int{2, 3}
	assertions.Equal(float64(2.5), findMedianSortedArrays(num3, num4))
}
