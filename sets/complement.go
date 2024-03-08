package sets

import "github.com/siliconbrain/go-seqs/seqs"

// ComplementOf returns the complement of the specified set.
func ComplementOf[T any](s SetOf[T]) SetOf[T] {
	return complementSet[T]{
		Set: s,
	}
}

// RelativeComplementOf returns the relative complement of set A in set B.
//
// This is also known as the set difference of B and A, denoted B \ A or B - A.
func RelativeComplementOf[T any](setA, setB SetOf[T]) SetOf[T] {
	return relativeComplementSet[SetOf[T], T]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableRelativeComplementOf returns the countable relative complement of set A in set B.
func CountableRelativeComplementOf[T any](setA, setB CountableSetOf[T]) CountableSetOf[T] {
	return countableRelativeComplementSet[T]{
		relativeComplementSet: relativeComplementSet[CountableSetOf[T], T]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type complementSet[T any] struct {
	Set SetOf[T]
}

func (s complementSet[T]) Contains(v T) bool {
	return !s.Set.Contains(v)
}

type relativeComplementSet[S SetOf[T], T any] struct {
	SetA S
	SetB S
}

func (s relativeComplementSet[_, T]) Contains(v T) bool {
	return s.SetB.Contains(v) && !s.SetA.Contains(v)
}

type countableRelativeComplementSet[T any] struct {
	relativeComplementSet[CountableSetOf[T], T]
}

func (s countableRelativeComplementSet[T]) AsSeq() seqs.Seq[T] {
	return seqs.Filter(s.SetB.AsSeq(), ComplementOf(s.SetA).Contains)
}
