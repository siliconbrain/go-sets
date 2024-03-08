package sets

import "github.com/siliconbrain/go-seqs/seqs"

// IntersectionOf returns the intersection of the two specified sets.
func IntersectionOf[T any](setA, setB SetOf[T]) SetOf[T] {
	cntSetA, okA := setA.(CountableSetOf[T])
	cntSetB, okB := setB.(CountableSetOf[T])
	if okA && okB {
		return CountableIntersectionOf(cntSetA, cntSetB)
	}
	return intersectionSet[SetOf[T], T]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableIntersectionOf returns the intersection of the two specified countable sets.
func CountableIntersectionOf[T any](setA, setB CountableSetOf[T]) CountableSetOf[T] {
	return countableIntersectionSet[T]{
		intersectionSet: intersectionSet[CountableSetOf[T], T]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type intersectionSet[S SetOf[T], T any] struct {
	SetA S
	SetB S
}

func (s intersectionSet[_, T]) Contains(v T) bool {
	return s.SetA.Contains(v) && s.SetB.Contains(v)
}

type countableIntersectionSet[T any] struct {
	intersectionSet[CountableSetOf[T], T]
}

func (s countableIntersectionSet[T]) ForEachUntil(fn func(T) bool) {
	smaller, larger := s.SetA, s.SetB

	cardS, hasCardS := QuickCardinalityOf(smaller)
	cardL, hasCardL := QuickCardinalityOf(larger)
	if hasCardL && hasCardS && cardL < cardS {
		// if our guess is verifiably wrong, swap them
		smaller, larger = larger, smaller
	}

	seqs.Filter(smaller, larger.Contains).ForEachUntil(fn)
}
