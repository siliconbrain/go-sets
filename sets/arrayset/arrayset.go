package arrayset

import (
	"fmt"
	"iter"
	"slices"

	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/siliconbrain/go-sets/sets/core"
)

func FromIter[Obj comparable](seq iter.Seq[Obj]) ArraySet[Obj] {
	return WithEq(equals[Obj]).FromIter(seq)
}

func FromSeq[Obj comparable](seq seqs.Seq[Obj]) ArraySet[Obj] {
	return WithEq(equals[Obj]).FromSeq(seq)
}

func FromSlice[Obj comparable](objs []Obj) ArraySet[Obj] {
	return WithEq(equals[Obj]).FromSlice(objs)
}

func FromValues[Obj comparable](objs ...Obj) ArraySet[Obj] {
	return WithEq(equals[Obj]).FromValues(objs...)
}

type ArraySet[Obj any] struct {
	Items []Obj
	Eq    func(Obj, Obj) bool
}

func (s ArraySet[_]) Cardinality() (core.Cardinal, bool) {
	c, err := core.CardinalFromInt(s.Len())
	return c, err == nil
}

func (s ArraySet[Obj]) Clone() ArraySet[Obj] {
	s.Items = slices.Clone(s.Items)
	return s
}

func (s ArraySet[Obj]) Contains(obj Obj) bool {
	return slices.ContainsFunc(s.Items, func(o Obj) bool {
		return s.Eq(o, obj)
	})
}

func (s *ArraySet[Obj]) Exclude(objs ...Obj) {
	if len(objs) == 0 {
		return
	}
	if s.Len() == 0 {
		return
	}
	s.ExcludeSet(s.WithItems(objs))
}

// ExcludeAll removes all objects from the set that are found in the specified sequence.
//
// NOTE: The current implementation is slower than calling Exclude(slices.Collect(seq)...) would be if converting the sequence to a slice is practical.
func (s *ArraySet[Obj]) ExcludeAll(seq iter.Seq[Obj]) {
	if s.Len() == 0 {
		return
	}
	for obj := range seq {
		s.Exclude(obj)
		if s.Len() == 0 {
			return
		}
	}
}

// ExcludeSeq removes all objects from the set that are found in the specified sequence.
//
// NOTE: The current implementation is slower than calling Exclude(seqs.ToSlice(seq)...) would be if converting the sequence to a slice is practical.
func (s *ArraySet[Obj]) ExcludeSeq(seq seqs.Seq[Obj]) {
	if seq, ok := seq.(seqs.Lener); ok && seq.Len() == 0 {
		return
	}
	s.ExcludeAll(seqs.ToIter(seq))
}

// ExcludeSet removes all items from the array set that are contained in the specified set.
func (s *ArraySet[Obj]) ExcludeSet(set core.SetOf[Obj]) {
	if s.Len() == 0 {
		return
	}
	if c, ok := core.QuickCardinalityOf(set); ok && c == 0 {
		return
	}
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

func (s ArraySet[Obj]) ForEachUntil(yield func(Obj) bool) {
	seqs.FromSlice(s.Items).ForEachUntil(yield)
}

func (s *ArraySet[Obj]) Include(objs ...Obj) {
	if len(objs) == 0 {
		return
	}
	s.IncludeAll(slices.Values(objs))
}

func (s *ArraySet[Obj]) IncludeAll(seq iter.Seq[Obj]) {
	for obj := range seq {
		if !s.Contains(obj) {
			s.Items = append(s.Items, obj)
		}
	}
}

func (s *ArraySet[Obj]) IncludeSeq(seq seqs.Seq[Obj]) {
	if seq, ok := seq.(seqs.Lener); ok && seq.Len() == 0 {
		return
	}
	s.IncludeAll(seqs.ToIter(seq))
}

func (s ArraySet[_]) Len() int {
	return len(s.Items)
}

func (s ArraySet[Obj]) Values(yield func(Obj) bool) {
	slices.Values(s.Items)(yield)
}

// WithEq returns a copy of the array set with its equality operator overridden with the specified function.
func (s ArraySet[Obj]) WithEq(eq func(Obj, Obj) bool) ArraySet[Obj] {
	if eq == nil {
		panic("eq must not be nil")
	}
	s.Eq = eq
	return s
}

// WithItems returns a copy of the array set with its items overridden with the specified items.
func (s ArraySet[Obj]) WithItems(items []Obj) ArraySet[Obj] {
	s.Items = uniqueInPlace(s.Eq, items)
	return s
}

var _ interface {
	core.CountableSetOf[any]
	core.Modifiable[any]
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
