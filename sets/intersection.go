package sets

import "github.com/siliconbrain/go-seqs/seqs"

// IntersectionOf returns the intersection of the two specified sets.
func IntersectionOf[SetA, SetB SetOf[Obj], Obj any](setA SetA, setB SetB) SetOf[Obj] {
	cntSetA, okA := any(setA).(CountableSetOf[Obj])
	cntSetB, okB := any(setB).(CountableSetOf[Obj])
	if okA && okB {
		return CountableIntersectionOf(cntSetA, cntSetB)
	}
	return intersectionSet[SetA, SetB, Obj]{
		SetA: setA,
		SetB: setB,
	}
}

// CountableIntersectionOf returns the intersection of the two specified countable sets.
func CountableIntersectionOf[SetA, SetB CountableSetOf[Obj], Obj any](setA SetA, setB SetB) CountableSetOf[Obj] {
	return countableIntersectionSet[SetA, SetB, Obj]{
		intersectionSet: intersectionSet[SetA, SetB, Obj]{
			SetA: setA,
			SetB: setB,
		},
	}
}

type intersectionSet[SetA, SetB SetOf[Obj], Obj any] struct {
	SetA SetA
	SetB SetB
}

func (s intersectionSet[_, _, Obj]) Contains(obj Obj) bool {
	return s.SetA.Contains(obj) && s.SetB.Contains(obj)
}

type countableIntersectionSet[SetA, SetB CountableSetOf[Obj], Obj any] struct {
	intersectionSet[SetA, SetB, Obj]
}

func (s countableIntersectionSet[_, _, Obj]) ForEachUntil(yield func(Obj) bool) {
	smaller, larger := CountableSetOf[Obj](s.SetA), CountableSetOf[Obj](s.SetB)

	cardS, hasCardS := QuickCardinalityOf(smaller)
	cardL, hasCardL := QuickCardinalityOf(larger)
	if hasCardL && hasCardS && cardL < cardS {
		// if our guess is verifiably wrong, swap them
		smaller, larger = larger, smaller
	}

	seqs.Filter(smaller, larger.Contains).ForEachUntil(yield)
}

func (s countableIntersectionSet[_, _, Obj]) Values(yield func(Obj) bool) {
	seqs.ToIter(s)(yield)
}
