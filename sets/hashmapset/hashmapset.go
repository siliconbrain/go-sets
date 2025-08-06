package hashmapset

import (
	"iter"
	"maps"
	"slices"

	"github.com/siliconbrain/go-seqs/mapseqs"
	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/siliconbrain/go-sets/sets/core"
)

// New returns an empty [HashMapSet] with the specified key extraction function.
func New[Obj any, Key comparable](getKey func(Obj) Key) HashMapSet[Obj, Key] {
	if getKey == nil {
		panic("getKey must not be nil")
	}
	return HashMapSet[Obj, Key]{
		getKey: getKey,
	}
}

// HashMapSet is a set of objects which are compared using a key derived from them.
type HashMapSet[Obj any, Key comparable] struct {
	hashmap map[Key]Obj
	getKey  func(Obj) Key
}

func (s HashMapSet[_, _]) Cardinality() (core.Cardinal, bool) {
	c, err := core.CardinalFromInt(s.Len())
	return c, err == nil
}

func (s HashMapSet[Obj, Key]) Clone() HashMapSet[Obj, Key] {
	return HashMapSet[Obj, Key]{
		hashmap: maps.Clone(s.hashmap),
		getKey:  s.getKey,
	}
}

func (s HashMapSet[Obj, _]) Contains(obj Obj) bool {
	_, res := s.hashmap[s.getKey(obj)]
	return res
}

func (s *HashMapSet[Obj, _]) Exclude(objs ...Obj) {
	if len(objs) == 0 {
		return
	}
	if len(s.hashmap) == 0 {
		return
	}
	s.ExcludeAll(slices.Values(objs))
}

func (s *HashMapSet[Obj, _]) ExcludeAll(seq iter.Seq[Obj]) {
	if len(s.hashmap) == 0 {
		return
	}
	seq(func(obj Obj) bool {
		delete(s.hashmap, s.getKey(obj))
		return len(s.hashmap) > 0
	})
}

func (s *HashMapSet[Obj, _]) ExcludeSeq(seq seqs.Seq[Obj]) {
	if len(s.hashmap) == 0 {
		return
	}
	if seq, ok := seq.(seqs.Lener); ok && seq.Len() == 0 {
		return
	}
	s.ExcludeAll(seqs.ToIter(seq))
}

func (s HashMapSet[Obj, _]) ForEachUntil(yield func(Obj) bool) {
	mapseqs.ValuesOf(s.hashmap).ForEachUntil(yield)
}

func (s HashMapSet[Obj, Key]) GetByKey(key Key) (Obj, bool) {
	obj, set := s.hashmap[key]
	return obj, set
}

func (s *HashMapSet[Obj, _]) Include(objs ...Obj) {
	if len(objs) == 0 {
		return
	}
	s.IncludeAll(slices.Values(objs))
}

func (s *HashMapSet[Obj, _]) IncludeAll(seq iter.Seq[Obj]) {
	s.ensureHashmap()
	seq(func(obj Obj) bool {
		s.hashmap[s.getKey(obj)] = obj
		return true
	})
}

func (s *HashMapSet[Obj, _]) IncludeSeq(seq seqs.Seq[Obj]) {
	if seq, ok := seq.(seqs.Lener); ok && seq.Len() == 0 {
		return
	}
	s.IncludeAll(seqs.ToIter(seq))
}

func (s HashMapSet[Obj, Key]) Len() int {
	return len(s.hashmap)
}

func (s HashMapSet[Obj, _]) Values(yield func(Obj) bool) {
	maps.Values(s.hashmap)(yield)
}

func (s *HashMapSet[Obj, Key]) ensureHashmap() {
	if s.hashmap == nil {
		s.hashmap = make(map[Key]Obj)
	}
}

var _ interface {
	core.CountableSetOf[any]
	core.Modifiable[any]
	seqs.FiniteSeq[any]
} = (*HashMapSet[any, any])(nil)
