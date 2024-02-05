package sets

import (
	"github.com/siliconbrain/go-mapseqs/mapseqs"
	"github.com/siliconbrain/go-seqs/seqs"
)

func HashSetFromSeq[T comparable](seq seqs.Seq[T]) (res HashSet[T]) {
	res.IncludeSeq(seq)
	return
}

func HashSetFromValues[T comparable](vs ...T) (res HashSet[T]) {
	res.Include(vs...)
	return
}

type HashSet[T comparable] map[T]struct{}

func (s HashSet[T]) AsSeq() seqs.Seq[T] {
	return mapseqs.KeysOf(s)
}

func (s HashSet[T]) Cardinality() int {
	return len(s)
}

func (s HashSet[T]) Contains(v T) (res bool) {
	_, res = s[v]
	return
}

func (s *HashSet[T]) Exclude(vs ...T) {
	if *s == nil {
		return
	}
	for _, v := range vs {
		delete(*s, v)
	}
}

func (s *HashSet[T]) ExcludeSeq(seq seqs.Seq[T]) {
	if *s == nil {
		return
	}
	seqs.ForEach(seq, func(v T) {
		delete(*s, v)
	})
}

func (s *HashSet[T]) Include(vs ...T) {
	if *s == nil {
		*s = make(HashSet[T])
	}
	for _, v := range vs {
		(*s)[v] = struct{}{}
	}
}

func (s *HashSet[T]) IncludeSeq(seq seqs.Seq[T]) {
	if *s == nil {
		*s = make(HashSet[T])
	}
	seqs.ForEach(seq, func(v T) {
		(*s)[v] = struct{}{}
	})
}

var _ interface {
	SetOf[any]
	Countable[any]
	Modifiable[any]
} = (*HashSet[any])(nil)
