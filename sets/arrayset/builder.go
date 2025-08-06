package arrayset

import (
	"iter"
	"slices"

	"github.com/siliconbrain/go-seqs/seqs"
)

func WithEq[Obj any](eq func(Obj, Obj) bool) Builder[Obj] {
	return Builder[Obj]{
		eq: eq,
	}
}

type Builder[Obj any] struct {
	eq func(Obj, Obj) bool
}

func (b Builder[Obj]) FromIter(seq iter.Seq[Obj]) ArraySet[Obj] {
	return b.FromValues(slices.Collect(seq)...)
}

func (b Builder[Obj]) FromSeq(seq seqs.Seq[Obj]) ArraySet[Obj] {
	return b.FromValues(seqs.ToSlice(seq)...)
}

func (b Builder[Obj]) FromSlice(objs []Obj) ArraySet[Obj] {
	return b.FromValues(slices.Clone(objs)...)
}

func (b Builder[Obj]) FromValues(objs ...Obj) ArraySet[Obj] {
	return ArraySet[Obj]{
		Items: uniqueInPlace(b.eq, objs),
		Eq:    b.eq,
	}
}

func uniqueInPlace[Obj any](eq func(Obj, Obj) bool, objs []Obj) []Obj {
	l := 0
	for i, obj := range objs {
		if !slices.ContainsFunc(objs[:l], func(o Obj) bool { return eq(obj, o) }) {
			if l != i { // only copy obj when the indices differ
				objs[l] = obj
			}
			l++
		}
	}
	return objs[:l]
}
