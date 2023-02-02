package solution1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	assertions := assert.New(t)

	s := "4193 with words"
	assertions.Equal(4193, myAtoi(s))

	s = "words and 987"
	assertions.Equal(0, myAtoi(s))

	s = ".1"
	assertions.Equal(0, myAtoi(s))

	s = "+-12"
	assertions.Equal(0, myAtoi(s))

	s = "    -42"
	assertions.Equal(-42, myAtoi(s))

	s = "9223372036854775808"
	assertions.Equal(2147483647, myAtoi(s))
}
