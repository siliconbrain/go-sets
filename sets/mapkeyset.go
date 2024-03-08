package sets

import (
	"github.com/siliconbrain/go-mapseqs/mapseqs"
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

func (s MapKeySet[M, K, V]) Contains(v K) bool {
	_, res := s.Map[v]
	return res
}

func (s MapKeySet[M, K, V]) ForEachUntil(fn func(K) bool) {
	mapseqs.KeysOf(s.Map).ForEachUntil(fn)
}

var _ CountableSetOf[any] = MapKeySet[map[any]any, any, any]{}
