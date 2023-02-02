package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {

	assertion := assert.New(t)
	s := "babad"
	assertion.Equal("bab", longestPalindrome(s))

	s = "cbbd"
	assertion.Equal("bb", longestPalindrome(s))

	s = "a"
	assertion.Equal("a", longestPalindrome(s))

	s = "ac"
	assertion.Equal("a", longestPalindrome(s))

	s = "ccc"
	assertion.Equal("ccc", longestPalindrome(s))
}
