package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	num1 := 121
	num2 := -121
	num3 := 10
	assertions := assert.New(t)

	assertions.True(isPalindrome(num1))
	assertions.False(isPalindrome(num2))
	assertions.False(isPalindrome(num3))
}
