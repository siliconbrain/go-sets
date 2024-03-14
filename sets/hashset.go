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

func (s *HashSet[T]) Exclude(vs ...T) {
	if s.Map == nil {
		return
	}
	for _, v := range vs {
		delete(s.Map, v)
	}
}

func (s *HashSet[T]) ExcludeSeq(seq seqs.Seq[T]) {
	if s.Map == nil {
		return
	}
	seqs.ForEach(seq, func(v T) {
		delete(s.Map, v)
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
} = (*HashSet[any])(nil)
