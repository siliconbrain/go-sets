package sets

import "github.com/siliconbrain/go-seqs/seqs"

// UnionOf returns the union of the two specified sets.
func UnionOf[SetA, SetB SetOf[Item], Item any](setA SetA, setB SetB) SetOf[Item] {
	cntSetA, okA := any(setA).(CountableSetOf[Item])
	cntSetB, okB := any(setB).(CountableSetOf[Item])
	if okA && okB {
		return CountableUnionOf(cntSetA, cntSetB)
	}
	return unionSet[SetA, SetB, Item]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableUnionOf returns the union of the two specified countable sets.
func CountableUnionOf[SetA, SetB CountableSetOf[Item], Item any](setA SetA, setB SetB) CountableSetOf[Item] {
	return countableUnionSet[SetA, SetB, Item]{
		unionSet: unionSet[SetA, SetB, Item]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type unionSet[SetA, SetB SetOf[Item], Item any] struct {
	SetA SetA
	SetB SetB
}

func (s unionSet[_, _, Item]) Contains(item Item) bool {
	return s.SetA.Contains(item) || s.SetB.Contains(item)
}

type countableUnionSet[SetA, SetB CountableSetOf[Item], Item any] struct {
	unionSet[SetA, SetB, Item]
}

func (s countableUnionSet[_, _, Item]) ForEachUntil(fn func(Item) bool) {
	seqs.Concat(s.SetA, seqs.Filter(s.SetB, ComplementOf(s.SetA).Contains)).ForEachUntil(fn)
}
