package core

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCardinal(t *testing.T) {
	t.Run("non-negative ints are representable as finite cardinals", func(t *testing.T) {
		for i := range intsBetweenZeroAndMaxInt {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				c, err := CardinalFromInt(i)
				require.NoError(t, err)
				assert.True(t, c.IsFinite())
			})
		}
	})

	t.Run("non-negative ints are recoverable from cardinals", func(t *testing.T) {
		for i := range intsBetweenZeroAndMaxInt {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				c, err := CardinalFromInt(i)
				require.NoError(t, err)
				ci, ok := c.AsInt()
				require.True(t, ok)
				assert.Equal(t, i, ci)
			})
		}
	})

	t.Run("large integer cardinal is finite", func(t *testing.T) {
		assert.True(t, CardinalLarge.IsFinite())
	})

	t.Run("aleph zero is not finite", func(t *testing.T) {
		assert.False(t, CardinalAleph0.IsFinite())
	})

	t.Run("cardinals are ordered", func(t *testing.T) {
		zero, _ := CardinalFromInt(0)
		fourTwo, _ := CardinalFromInt(42)
		maxInt, _ := CardinalFromInt(math.MaxInt)
		assert.Less(t, zero, fourTwo)
		assert.Less(t, fourTwo, maxInt)
		assert.Less(t, maxInt, CardinalLarge)
		assert.Less(t, CardinalLarge, CardinalAleph0)
	})
}

// intsBetweenZeroAndMaxInt yields ints between 0 and math.MaxInt, inclusive.
func intsBetweenZeroAndMaxInt(yield func(int) bool) {
	if !yield(0) {
		return
	}
	a, b := uint(1), uint(2)
	for a < math.MaxInt {
		if !yield(int(a)) {
			return
		}
		a, b = b, a+b
	}
	_ = yield(math.MaxInt)
}
