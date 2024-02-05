package sets

import "github.com/siliconbrain/go-seqs/seqs"

// SetOf defines the minimal interface of a set of T.
type SetOf[T any] interface {
	Contains(T) bool
}

// CountableSetOf defines the interface of a countable set of T.
type CountableSetOf[T any] interface {
	SetOf[T]
	Countable[T]
}

// Countable defines something whose elements can be listed, albeit not necessarily finite.
type Countable[T any] interface {
	AsSeq() seqs.Seq[T]
}

// Modifiable defines how a set can be modified.
type Modifiable[T any] interface {
	Exclude(...T)
	Include(...T)
}

// AreEqual returns true if the specified sets contain the same elements.
func AreEqual[SetA, SetB CountableSetOf[T], T any](setA SetA, setB SetB) bool {
	if setA, ok := any(setA).(interface{ EqualTo(SetB) bool }); ok {
		return setA.EqualTo(setB)
	}
	if setB, ok := any(setB).(interface{ EqualTo(SetA) bool }); ok {
		return setB.EqualTo(setA)
	}
	cardA, hasCardA := QuickCardinalityOf(setA)
	cardB, hasCardB := QuickCardinalityOf(setB)
	if hasCardA && hasCardB {
		return cardA == cardB && seqs.All(setA.AsSeq(), setB.Contains)
	}
	return seqs.All(setA.AsSeq(), setB.Contains) && seqs.All(setB.AsSeq(), setA.Contains)
}

// CardinalityOf returns the cardinality of the set.
//
// Calling it on an infinite set will result in an infinite loop.
func CardinalityOf[T any](s CountableSetOf[T]) int {
	if c, ok := QuickCardinalityOf(s); ok {
		return c
	}
	return seqs.Reduce(s.AsSeq(), 0, func(c int, _ T) int { return c + 1 })
}

// ContainsAnyOf returns true if s contains any of the specified values
func ContainsAnyOf[T any](s CountableSetOf[T], vs ...T) bool {
	for _, v := range vs {
		if s.Contains(v) {
			return true
		}
	}
	return false
}

// QuickCardinalityOf returns the cardinality of the set if it can be determined without counting its elements.
func QuickCardinalityOf[T any](s CountableSetOf[T]) (int, bool) {
	if s, ok := s.(interface{ Cardinality() int }); ok {
		return s.Cardinality(), true
	}
	if lener, ok := s.AsSeq().(seqs.Lener); ok {
		return lener.Len(), true
	}
	return 0, false
}

// ToSlice returns the elements of the specified set as a slice.
//
// The order of elements is undefined.
func ToSlice[T any](s CountableSetOf[T]) []T {
	return seqs.ToSlice(s.AsSeq())
}
