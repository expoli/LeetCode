package solution1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestName(t *testing.T) {
	assertion := assert.New(t)

	x := 123
	assertion.Equal(321, reverse(x))

	x = -123
	assertion.Equal(-321, reverse(x))

	x = 120
	assertion.Equal(21, reverse(x))

	x = 0
	assertion.Equal(0, reverse(x))

	x = 1534236469
	assertion.Equal(0, reverse(x))

	fmt.Println(9646324351 > math.MaxInt32)
}
