package sets

import "github.com/siliconbrain/go-seqs/seqs"

// UnionOf returns the union of the two specified sets.
func UnionOf[T any](setA, setB SetOf[T]) SetOf[T] {
	cntSetA, okA := setA.(CountableSetOf[T])
	cntSetB, okB := setB.(CountableSetOf[T])
	if okA && okB {
		return CountableUnionOf(cntSetA, cntSetB)
	}
	return unionSet[SetOf[T], T]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableUnionOf returns the union of the two specified countable sets.
func CountableUnionOf[T any](setA, setB CountableSetOf[T]) CountableSetOf[T] {
	return countableUnionSet[T]{
		unionSet: unionSet[CountableSetOf[T], T]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type unionSet[S SetOf[T], T any] struct {
	SetA S
	SetB S
}

func (s unionSet[_, T]) Contains(v T) bool {
	return s.SetA.Contains(v) || s.SetB.Contains(v)
}

type countableUnionSet[T any] struct {
	unionSet[CountableSetOf[T], T]
}

func (s countableUnionSet[T]) ForEachUntil(fn func(T) bool) {
	seqs.Concat(s.SetA, seqs.Filter(s.SetB, ComplementOf(s.SetA).Contains)).ForEachUntil(fn)
}
