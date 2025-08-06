package sets

import "github.com/siliconbrain/go-seqs/seqs"

// ComplementOf returns the complement of the specified set.
func ComplementOf[Set SetOf[Obj], Obj any](set Set) SetOf[Obj] {
	return complementSet[Set, Obj]{
		Set: set,
	}
}

// RelativeComplementOf returns the relative complement of set A in set B.
//
// This is also known as the set difference of B and A, denoted B \ A or B - A.
func RelativeComplementOf[SetA, SetB SetOf[Obj], Obj any](setA SetA, setB SetB) SetOf[Obj] {
	return relativeComplementSet[SetA, SetB, Obj]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableRelativeComplementOf returns the countable relative complement of set A in set B.
func CountableRelativeComplementOf[SetA, SetB CountableSetOf[Obj], Obj any](setA SetA, setB SetB) CountableSetOf[Obj] {
	return countableRelativeComplementSet[SetA, SetB, Obj]{
		relativeComplementSet: relativeComplementSet[SetA, SetB, Obj]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type complementSet[Set SetOf[Obj], Obj any] struct {
	Set Set
}

func (s complementSet[_, Obj]) Contains(obj Obj) bool {
	return !s.Set.Contains(obj)
}

type relativeComplementSet[SetA, SetB SetOf[Obj], Obj any] struct {
	SetA SetA
	SetB SetB
}

func (s relativeComplementSet[_, _, Obj]) Contains(obj Obj) bool {
	return s.SetB.Contains(obj) && !s.SetA.Contains(obj)
}

type countableRelativeComplementSet[SetA, SetB CountableSetOf[Obj], Obj any] struct {
	relativeComplementSet[SetA, SetB, Obj]
}

func (s countableRelativeComplementSet[_, _, Obj]) ForEachUntil(yield func(Obj) bool) {
	seqs.Filter(s.SetB, ComplementOf(s.SetA).Contains).ForEachUntil(yield)
}

func (s countableRelativeComplementSet[_, _, Obj]) Values(yield func(Obj) bool) {
	seqs.ToIter(s)(yield)
}
