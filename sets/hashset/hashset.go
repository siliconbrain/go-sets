package hashset

import (
	"iter"
	"slices"

	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/siliconbrain/go-sets/sets/core"
	"github.com/siliconbrain/go-sets/sets/mapkeyset"
)

func FromIter[Obj comparable](seq iter.Seq[Obj]) (res HashSet[Obj]) {
	res.IncludeAll(seq)
	return
}

func FromSeq[Obj comparable](seq seqs.Seq[Obj]) (res HashSet[Obj]) {
	res.IncludeSeq(seq)
	return
}

func FromSlice[Obj comparable](objs []Obj) (res HashSet[Obj]) {
	res.Include(objs...)
	return
}

func FromValues[Obj comparable](objs ...Obj) (res HashSet[Obj]) {
	return FromSlice(objs)
}

type HashSet[Obj comparable] struct {
	mapkeyset.MapKeySet[map[Obj]struct{}, Obj, struct{}]
}

func (s HashSet[Obj]) Clone() HashSet[Obj] {
	return HashSet[Obj]{
		MapKeySet: s.MapKeySet.Clone(),
	}
}

func (s *HashSet[Obj]) Exclude(objs ...Obj) {
	if len(objs) == 0 {
		return
	}
	s.ExcludeAll(slices.Values(objs))
}

func (s *HashSet[Obj]) ExcludeAll(seq iter.Seq[Obj]) {
	if s.Len() == 0 {
		return
	}
	for obj := range seq {
		delete(s.Map, obj)
		if s.Len() == 0 {
			return
		}
	}
}

func (s *HashSet[Obj]) ExcludeSeq(seq seqs.Seq[Obj]) {
	if seq, ok := seq.(seqs.Lener); ok && seq.Len() == 0 {
		return
	}
	s.ExcludeAll(seqs.ToIter(seq))
}

func (s *HashSet[Obj]) ExcludeSet(set core.SetOf[Obj]) {
	if s.Len() == 0 {
		return
	}
	if c, ok := core.QuickCardinalityOf(set); ok && c == 0 {
		return
	}
	if set, ok := set.(core.CountableSetOf[Obj]); ok {
		s.ExcludeSeq(set)
		return
	}
	for item := range s.Map {
		if set.Contains(item) {
			delete(s.Map, item)
			if s.Len() == 0 {
				return
			}
		}
	}
}

func (s *HashSet[Obj]) Include(objs ...Obj) {
	if len(objs) == 0 {
		return
	}
	s.IncludeAll(slices.Values(objs))
}

func (s *HashSet[Obj]) IncludeAll(seq iter.Seq[Obj]) {
	s.ensureMap()
	for obj := range seq {
		s.Map[obj] = struct{}{}
	}
}

func (s *HashSet[T]) IncludeSeq(seq seqs.Seq[T]) {
	if seq, ok := seq.(seqs.Lener); ok && seq.Len() == 0 {
		return
	}
	s.IncludeAll(seqs.ToIter(seq))
}

func (s *HashSet[Obj]) ensureMap() {
	if s.Map == nil {
		s.Map = make(map[Obj]struct{})
	}
}

var _ interface {
	core.CountableSetOf[any]
	core.Modifiable[any]
	seqs.FiniteSeq[any]
} = (*HashSet[any])(nil)
