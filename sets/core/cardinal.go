package core

import (
	"errors"
	"math"
)

// Cardinal represents a cardinal number.
type Cardinal uint

// AsInt returns the cardinal as an int (and true) when it is representable as a non-negative integer.
// Otherwise, it returns 0 and false.
func (c Cardinal) AsInt() (int, bool) {
	if c > math.MaxInt {
		return 0, false
	}
	return int(c), true
}

// IsFinite returns true when the cardinal is finite (i.e. a natural number).
// It does NOT mean that it is representable as an int or
func (c Cardinal) IsFinite() bool {
	return c < CardinalAleph0
}

// CardinalLarge represents a finite cardinal that is too large to be represented by a non-negative int.
const CardinalLarge Cardinal = math.MaxInt + 1

// CardinalAleph0 represents the cardinality of the natural numbers, the smallest infinite cardinal, commonly denoted by ℵ₀.
const CardinalAleph0 Cardinal = CardinalLarge + 1

// CardinalFromInt turns a non-negative int into a [Cardinal].
// It returns an error if the int is negative.
func CardinalFromInt(i int) (Cardinal, error) {
	if i < 0 {
		return 0, errors.New("cardinal must be non-negative")
	}
	return Cardinal(i), nil
}
