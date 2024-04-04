package arrayset

import "github.com/siliconbrain/go-seqs/seqs"

func WithEq[T any](eq func(T, T) bool) Builder[T] {
	return Builder[T]{
		eq: eq,
	}
}

type Builder[T any] struct {
	eq func(T, T) bool
}

func (b Builder[T]) FromSeq(seq seqs.Seq[T]) ArraySet[T] {
	return b.FromSlice(seqs.ToSlice(seq))
}

func (b Builder[T]) FromSlice(vs []T) ArraySet[T] {
	return ArraySet[T]{
		Items: vs,
		Eq:    b.eq,
	}
}

func (b Builder[T]) FromValues(vs ...T) ArraySet[T] {
	return b.FromSlice(vs)
}
