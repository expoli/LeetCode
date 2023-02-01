package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3

	a := assert.New(t)
	a.Equal("PAHNAPLSIIGYIR", convert(s, numRows))

	s = "PAYPALISHIRING"
	numRows = 4
	a.Equal("PINALSIGYAHRPI", convert(s, numRows))

	s = "A"
	numRows = 1
	a.Equal("A", convert(s, numRows))

	s = "AB"
	numRows = 1
	a.Equal("AB", convert(s, numRows))
}
