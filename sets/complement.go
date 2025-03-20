package sets

import "github.com/siliconbrain/go-seqs/seqs"

// ComplementOf returns the complement of the specified set.
func ComplementOf[Set SetOf[Item], Item any](s Set) SetOf[Item] {
	return complementSet[Set, Item]{
		Set: s,
	}
}

// RelativeComplementOf returns the relative complement of set A in set B.
//
// This is also known as the set difference of B and A, denoted B \ A or B - A.
func RelativeComplementOf[SetA, SetB SetOf[Item], Item any](setA SetA, setB SetB) SetOf[Item] {
	return relativeComplementSet[SetA, SetB, Item]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableRelativeComplementOf returns the countable relative complement of set A in set B.
func CountableRelativeComplementOf[SetA, SetB CountableSetOf[Item], Item any](setA SetA, setB SetB) CountableSetOf[Item] {
	return countableRelativeComplementSet[SetA, SetB, Item]{
		relativeComplementSet: relativeComplementSet[SetA, SetB, Item]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type complementSet[Set SetOf[Item], Item any] struct {
	Set Set
}

func (s complementSet[_, Item]) Contains(item Item) bool {
	return !s.Set.Contains(item)
}

type relativeComplementSet[SetA, SetB SetOf[Item], Item any] struct {
	SetA SetA
	SetB SetB
}

func (s relativeComplementSet[_, _, Item]) Contains(item Item) bool {
	return s.SetB.Contains(item) && !s.SetA.Contains(item)
}

type countableRelativeComplementSet[SetA, SetB CountableSetOf[Item], Item any] struct {
	relativeComplementSet[SetA, SetB, Item]
}

func (s countableRelativeComplementSet[_, _, Item]) ForEachUntil(fn func(Item) bool) {
	seqs.Filter(s.SetB, ComplementOf(s.SetA).Contains).ForEachUntil(fn)
}
