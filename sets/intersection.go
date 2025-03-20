package sets

import "github.com/siliconbrain/go-seqs/seqs"

// IntersectionOf returns the intersection of the two specified sets.
func IntersectionOf[SetA, SetB SetOf[Item], Item any](setA SetA, setB SetB) SetOf[Item] {
	cntSetA, okA := any(setA).(CountableSetOf[Item])
	cntSetB, okB := any(setB).(CountableSetOf[Item])
	if okA && okB {
		return CountableIntersectionOf(cntSetA, cntSetB)
	}
	return intersectionSet[SetA, SetB, Item]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableIntersectionOf returns the intersection of the two specified countable sets.
func CountableIntersectionOf[SetA, SetB CountableSetOf[Item], Item any](setA SetA, setB SetB) CountableSetOf[Item] {
	return countableIntersectionSet[SetA, SetB, Item]{
		intersectionSet: intersectionSet[SetA, SetB, Item]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type intersectionSet[SetA, SetB SetOf[Item], Item any] struct {
	SetA SetA
	SetB SetB
}

func (s intersectionSet[_, _, Item]) Contains(item Item) bool {
	return s.SetA.Contains(item) && s.SetB.Contains(item)
}

type countableIntersectionSet[SetA, SetB CountableSetOf[Item], Item any] struct {
	intersectionSet[SetA, SetB, Item]
}

func (s countableIntersectionSet[_, _, Item]) ForEachUntil(fn func(Item) bool) {
	smaller, larger := CountableSetOf[Item](s.SetA), CountableSetOf[Item](s.SetB)

	cardS, hasCardS := QuickCardinalityOf(smaller)
	cardL, hasCardL := QuickCardinalityOf(larger)
	if hasCardL && hasCardS && cardL < cardS {
		// if our guess is verifiably wrong, swap them
		smaller, larger = larger, smaller
	}

	seqs.Filter(smaller, larger.Contains).ForEachUntil(fn)
}
