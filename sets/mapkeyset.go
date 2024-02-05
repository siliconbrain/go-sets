package sets

import (
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

func (s MapKeySet[M, K, V]) AsSeq() seqs.Seq[K] {
	return mapseqs.KeysOf(s.Map)
}

func (s MapKeySet[M, K, V]) Cardinality() int {
	return len(s.Map)
}

func (s MapKeySet[M, K, V]) Contains(v K) bool {
	_, res := s.Map[v]
	return res
}

var _ CountableSetOf[any] = MapKeySet[map[any]any, any, any]{}
