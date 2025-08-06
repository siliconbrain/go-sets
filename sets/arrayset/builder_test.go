package arrayset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueInPlace(t *testing.T) {
	assert.Equal(t, ([]int)(nil), uniqueInPlace(equals[int], ([]int)(nil)))
	assert.Equal(t, []int{}, uniqueInPlace(equals[int], []int{}))
	assert.Equal(t, []int{1, 4, 3, 5, 6, 8, 0, 9}, uniqueInPlace(equals[int], []int{1, 4, 3, 5, 3, 3, 5, 6, 8, 0, 1, 9}))
}
