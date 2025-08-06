package byteset

import (
	"iter"
	"math/bits"
	"slices"

	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/siliconbrain/go-sets/sets/core"
)

// FromIter creates a [ByteSet] from the specified [iter.Seq] of bytes.
func FromIter(bytes iter.Seq[byte]) (set ByteSet) {
	for b := range bytes {
		offset, mask := offsetAndMask(b)
		set.bits[offset] |= mask
	}
	return
}

// FromSeq creates a [ByteSet] from the specified [seqs.Seq] of bytes.
func FromSeq(bytes seqs.Seq[byte]) ByteSet {
	return FromIter(seqs.ToIter(bytes))
}

// FromValues creates a [ByteSet] of the specified bytes.
//
// Example:
//
//	s := ByteSetFromValues('a', '!', 42)
//	// s will now look like this:
//	//          63                  42:'*'   33:'!'                             0
//	//           |                    |        |                                |
//	// bits[0]   0000000000000000000001000000001000000000000000000000000000000000
//	//         127                           97:'a'                             64
//	//           |                             |                                |
//	// bits[1]   0000000000000000000000000000001000000000000000000000000000000000
//	//         191                                                              128
//	//           |                                                              |
//	// bits[2]   0000000000000000000000000000000000000000000000000000000000000000
//	//         255                                                              192
//	//           |                                                              |
//	// bits[3]   0000000000000000000000000000000000000000000000000000000000000000
//	assert(s.Contains('*'))
//	assert(!s.Contains('A'))
//	assert(s.Len() == 3)
//	assert(slices.Collect(s.Values) == []byte{33, 42, 97})
func FromValues(bytes ...byte) ByteSet {
	return FromIter(slices.Values(bytes))
}

// ByteSet is a set of bytes (8 bit values).
//
// Internally, set membership is stored by a 256 bit bit array, each bit corresponding to a byte value.
// If a bit is set, the byte equal to its index is part of the set.
type ByteSet struct {
	bits [4]uint64 // 256 bits, one for each 1 byte value
}

func (set ByteSet) Cardinality() (core.Cardinal, bool) {
	c, err := core.CardinalFromInt(set.Len())
	return c, err == nil
}

// Contains return whether the [ByteSet] contains the specified byte
func (set ByteSet) Contains(b byte) bool {
	offset, mask := offsetAndMask(b)
	return set.bits[offset]&mask != 0
}

func (set ByteSet) ForEachUntil(yield func(byte) bool) {
	seqs.FromIter(set.Values).ForEachUntil(yield)
}

func (set ByteSet) Len() (n int) {
	for o := range set.bits {
		n += bits.OnesCount64(set.bits[o])
	}
	return
}

func (set ByteSet) Values(yield func(byte) bool) {
	for i := range 256 {
		if b := byte(i); set.Contains(b) {
			if !yield(b) {
				return
			}
		}
	}
}

// Complement returns the complement of the [ByteSet].
func Complement(set ByteSet) ByteSet {
	for offset, mask := range set.bits {
		set.bits[offset] = ^mask
	}
	return set
}

// Intersection returns the intersection of the specified [ByteSet]s.
// NOTE: Returns a [ByteSet] matching every byte if sets is empty.
func Intersection(sets ...ByteSet) (set ByteSet) {
	// init set to all 1's
	for i := range len(set.bits) {
		set.bits[i] = ^uint64(0)
	}
	for s := range slices.Values(sets) {
		for offset, mask := range s.bits {
			set.bits[offset] &= mask
		}
	}
	return
}

// RelativeComplement returns the relative complement of set A in set B.
//
// This is also known as the set difference of B and A, denoted B \ A or B - A.
func RelativeComplement(setA, setB ByteSet) (set ByteSet) {
	for offset := range set.bits {
		set.bits[offset] = setB.bits[offset] & ^setA.bits[offset]
	}
	return
}

// Union returns the union of the specified [ByteSet]s.
func Union(sets ...ByteSet) (set ByteSet) {
	for s := range slices.Values(sets) {
		for offset, mask := range s.bits {
			set.bits[offset] |= mask
		}
	}
	return
}

// offsetAndMask returns the index of the quad-word for which the mask is to be applied
func offsetAndMask(b byte) (offset byte, mask uint64) {
	return b / 64, uint64(1) << (b % 64)
}
