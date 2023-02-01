package golang

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	s := "babad"
	fmt.Println(longestPalindrome(s))

	s = "cbbd"
	fmt.Println(longestPalindrome(s))

	s = "a"
	fmt.Println(longestPalindrome(s))

	s = "ac"
	fmt.Println(longestPalindrome(s))
}
