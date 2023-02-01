package solution2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	str1 := "abcabcbb"
	str2 := "bbbbb"
	str3 := "pwwkew"
	str4 := " "
	str5 := "dvdf"
	str6 := "anviaj"

	assertions := assert.New(t)

	assertions.Equal(3, lengthOfLongestSubstring(str1))
	assertions.Equal(1, lengthOfLongestSubstring(str2))
	assertions.Equal(3, lengthOfLongestSubstring(str3))
	assertions.Equal(1, lengthOfLongestSubstring(str4))
	assertions.Equal(3, lengthOfLongestSubstring(str5))
	assertions.Equal(5, lengthOfLongestSubstring(str6))
}
