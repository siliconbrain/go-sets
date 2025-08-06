package hashmapset

import (
	"iter"

	"github.com/siliconbrain/go-seqs/seqs"
)

func WithGetKey[Obj any, Key comparable](getKey func(Obj) Key) Builder[Obj, Key] {
	if getKey == nil {
		panic("getKey must not be nil")
	}
	return Builder[Obj, Key]{
		getKey: getKey,
	}
}

type Builder[Obj any, Key comparable] struct {
	getKey func(Obj) Key
}

func (b Builder[Obj, Key]) FromIter(seq iter.Seq[Obj]) HashMapSet[Obj, Key] {
	res := New(b.getKey)
	res.IncludeAll(seq)
	return res
}

func (b Builder[Obj, Key]) FromSeq(seq seqs.Seq[Obj]) HashMapSet[Obj, Key] {
	res := New(b.getKey)
	res.IncludeSeq(seq)
	return res
}

func (b Builder[Obj, Key]) FromSlice(objs []Obj) HashMapSet[Obj, Key] {
	res := New(b.getKey)
	res.Include(objs...)
	return res
}

func (b Builder[Obj, Key]) FromValues(objs ...Obj) HashMapSet[Obj, Key] {
	res := New(b.getKey)
	res.Include(objs...)
	return res
}
