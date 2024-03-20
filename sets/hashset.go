package sets

import (
	"github.com/siliconbrain/go-seqs/seqs"
)

func HashSetFromSeq[T comparable](seq seqs.Seq[T]) (res HashSet[T]) {
	res.IncludeSeq(seq)
	return
}

func HashSetFromSlice[T comparable](vs []T) (res HashSet[T]) {
	res.Include(vs...)
	return
}

func HashSetFromValues[T comparable](vs ...T) (res HashSet[T]) {
	return HashSetFromSlice(vs)
}

type HashSet[T comparable] struct {
	MapKeySet[map[T]struct{}, T, struct{}]
}

func (s HashSet[T]) Clone() HashSet[T] {
	return HashSet[T]{
		MapKeySet: s.MapKeySet.Clone(),
	}
}

func (s *HashSet[T]) Exclude(vs ...T) {
	if len(s.Map) == 0 {
		return
	}
	for _, v := range vs {
		delete(s.Map, v)
		if len(s.Map) == 0 {
			return
		}
	}
}

func (s *HashSet[T]) ExcludeSeq(seq seqs.Seq[T]) {
	if len(s.Map) == 0 {
		return
	}
	seq.ForEachUntil(func(v T) bool {
		delete(s.Map, v)
		return len(s.Map) == 0
	})
}

func (s *HashSet[T]) Include(vs ...T) {
	if s.Map == nil {
		s.Map = make(map[T]struct{})
	}
	for _, v := range vs {
		s.Map[v] = struct{}{}
	}
}

func (s *HashSet[T]) IncludeSeq(seq seqs.Seq[T]) {
	if s.Map == nil {
		s.Map = make(map[T]struct{})
	}
	seqs.ForEach(seq, func(v T) {
		s.Map[v] = struct{}{}
	})
}

var _ interface {
	CountableSetOf[any]
	Modifiable[any]
	seqs.FiniteSeq[any]
} = (*HashSet[any])(nil)
