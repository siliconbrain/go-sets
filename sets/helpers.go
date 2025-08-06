package sets

import (
	"iter"
	"slices"

	"github.com/siliconbrain/go-seqs/seqs"
)

// AreEqual returns true if the specified sets contain the same objects.
func AreEqual[SetA, SetB CountableSetOf[Obj], Obj any](setA SetA, setB SetB) bool {
	if setA, ok := any(setA).(interface{ EqualTo(SetB) bool }); ok {
		return setA.EqualTo(setB)
	}
	if setB, ok := any(setB).(interface{ EqualTo(SetA) bool }); ok {
		return setB.EqualTo(setA)
	}
	cardA, hasCardA := QuickCardinalityOf(setA)
	cardB, hasCardB := QuickCardinalityOf(setB)
	if hasCardA && hasCardB {
		return cardA == cardB && seqs.All(setA, setB.Contains)
	}
	return seqs.All(setA, setB.Contains) && seqs.All(setB, setA.Contains)
}

// CardinalityOf returns the cardinality of the set if it can be determined.
//
// Calling it on an infinite set might result in an infinite loop.
func CardinalityOf[Obj any](set SetOf[Obj]) (Cardinal, bool) {
	if c, ok := QuickCardinalityOf(set); ok {
		return c, ok
	}
	if set, ok := set.(CountableSetOf[Obj]); ok {
		c, err := CardinalFromInt(seqs.Count(set))
		return c, err == nil
	}
	return 0, false
}

// ContainsAllOf returns true if set contains all of the specified objects.
func ContainsAllOf[Obj any](set SetOf[Obj], objs ...Obj) bool {
	return ContainsAllOfIter(set, slices.Values(objs))
}

// ContainsAllOfIter return true if set contains all of the objects in seq.
func ContainsAllOfIter[Obj any](set SetOf[Obj], seq iter.Seq[Obj]) bool {
	for obj := range seq {
		if !set.Contains(obj) {
			return false
		}
	}
	return true
}

// ContainsAnyOf returns true if set contains any of the specified objects.
func ContainsAnyOf[Obj any](set SetOf[Obj], objs ...Obj) bool {
	return ContainsAnyOfIter(set, slices.Values(objs))
}

// ContainsAnyOfIter returns true if set contains any of the objects in seq.
func ContainsAnyOfIter[Obj any](set SetOf[Obj], seq iter.Seq[Obj]) bool {
	for obj := range seq {
		if set.Contains(obj) {
			return true
		}
	}
	return false
}

// IsEmpty returns (true, true) when the set is empty, (false, true) when the set is not empty, and (false, false) when it is unknown.
func IsEmpty[Obj any](set SetOf[Obj]) (empty bool, known bool) {
	if fin, ok := set.(interface{ Len() int }); ok {
		return fin.Len() == 0, true
	}
	if card, ok := QuickCardinalityOf(set); ok {
		return card == 0, true
	}
	if seq, ok := set.(seqs.Seq[Obj]); ok {
		return seqs.IsEmpty(seq), true
	}
	return false, false
}
