package arrayset

import (
	"fmt"
	"slices"

	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/siliconbrain/go-sets/sets"
)

func FromSeq[T comparable](seq seqs.Seq[T]) ArraySet[T] {
	return WithEq(equals[T]).FromSeq(seq)
}

func FromSlice[T comparable](vs []T) ArraySet[T] {
	return WithEq(equals[T]).FromSlice(vs)
}

func FromValues[T comparable](vs ...T) ArraySet[T] {
	return WithEq(equals[T]).FromValues(vs...)
}

type ArraySet[T any] struct {
	Items []T
	Eq    func(T, T) bool
}

func (s ArraySet[T]) Cardinality() int {
	return len(s.Items)
}

func (s ArraySet[T]) Clone() ArraySet[T] {
	s.Items = slices.Clone(s.Items)
	return s
}

func (s ArraySet[T]) Contains(v T) bool {
	return slices.ContainsFunc(s.Items, func(i T) bool {
		return s.Eq(i, v)
	})
}

func (s *ArraySet[T]) Exclude(vs ...T) {
	if s.Cardinality() == 0 {
		return
	}
	s.ExcludeSet(s.WithItems(vs))
}

// ExcludeSeq removes all items from the set that are found in the specified sequence.
//
// NOTE: The current implementation is slower than calling Exclude(seqs.ToSlice(seq)) would be if converting the sequence to a slice is practical.
func (s *ArraySet[T]) ExcludeSeq(seq seqs.Seq[T]) {
	if s.Cardinality() == 0 {
		return
	}
	seq.ForEachUntil(func(v T) bool {
		s.Exclude(v)
		return s.Cardinality() == 0
	})
}

// ExcludeSet removes all items from the array set that are contained in the specified set.
func (s *ArraySet[T]) ExcludeSet(set sets.SetOf[T]) {
	for index := len(s.Items) - 1; index >= 0; index-- {
		if set.Contains(s.Items[index]) {
			s.Items = sliceRemoveAtUnordered(s.Items, index)
		}
		// the item at index is either
		// * determined to be a kept,
		// * has been removed (if at end), or
		// * replaced by a value that is a known keeper
		// thus, index can be decremented
	}
}

func (s ArraySet[T]) ForEachUntil(fn func(T) bool) {
	seqs.FromSlice(s.Items).ForEachUntil(fn)
}

func (s *ArraySet[T]) Include(vs ...T) {
	if len(vs) == 0 {
		return
	}
	s.IncludeSeq(seqs.FromSlice(vs))
}

func (s *ArraySet[T]) IncludeSeq(seq seqs.Seq[T]) {
	seqs.ForEach(seq, func(v T) {
		if !s.Contains(v) {
			s.Items = append(s.Items, v)
		}
	})
}

func (s ArraySet[T]) Len() int {
	return s.Cardinality()
}

// WithEq returns a copy of the array set with its equality operator overridden with the specified function.
func (s ArraySet[T]) WithEq(eq func(T, T) bool) ArraySet[T] {
	if eq == nil {
		panic("eq must not be nil")
	}
	s.Eq = eq
	return s
}

// WithItems returns a copy of the array set with its items overridden with the specified items.
func (s ArraySet[T]) WithItems(items []T) ArraySet[T] {
	s.Items = items
	return s
}

var _ interface {
	sets.CountableSetOf[any]
	sets.Modifiable[any]
	seqs.FiniteSeq[any]
} = (*ArraySet[any])(nil)

func equals[T comparable](a, b T) bool {
	return a == b
}

func sliceRemoveAtUnordered[T any](slice []T, index int) []T {
	lastIndex := len(slice) - 1
	if index < 0 || index > lastIndex {
		panic(fmt.Errorf("index [%d] out of range with length %d", index, len(slice)))
	}
	if index != lastIndex { // only fill hole with last item if the hole itself is not the last item
		slice[index] = slice[lastIndex]
	}
	return slice[:lastIndex] // drop last item
}
