package sets

import (
	"maps"

	"github.com/siliconbrain/go-mapseqs/mapseqs"
	"github.com/siliconbrain/go-seqs/seqs"
)

func MapKeySetFrom[M ~map[K]V, K comparable, V any](m M) MapKeySet[M, K, V] {
	return MapKeySet[M, K, V]{
		Map: m,
	}
}

type MapKeySet[M ~map[K]V, K comparable, V any] struct {
	Map M
}

func (s MapKeySet[M, K, V]) Cardinality() int {
	return len(s.Map)
}

func (s MapKeySet[M, K, V]) Clone() MapKeySet[M, K, V] {
	return MapKeySetFrom(maps.Clone(s.Map))
}

func (s MapKeySet[M, K, V]) Contains(v K) bool {
	_, res := s.Map[v]
	return res
}

func (s MapKeySet[M, K, V]) ForEachUntil(fn func(K) bool) {
	mapseqs.KeysOf(s.Map).ForEachUntil(fn)
}

func (s MapKeySet[M, K, V]) Len() int {
	return len(s.Map)
}

var _ interface {
	CountableSetOf[any]
	seqs.FiniteSeq[any]
} = MapKeySet[map[any]any, any, any]{}
