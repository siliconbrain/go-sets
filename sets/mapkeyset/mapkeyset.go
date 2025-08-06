package mapkeyset

import (
	"maps"

	"github.com/siliconbrain/go-seqs/mapseqs"
	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/siliconbrain/go-sets/sets/core"
)

// FromMap returns a new [MapKeySet] based on the specified map-like value.
func FromMap[Map ~map[Key]Val, Key comparable, Val any](m Map) MapKeySet[Map, Key, Val] {
	return MapKeySet[Map, Key, Val]{
		Map: m,
	}
}

// MapKeySet implements a set using the keys of a map-like value.
type MapKeySet[Map ~map[Key]Val, Key comparable, Val any] struct {
	Map Map
}

func (s MapKeySet[_, _, _]) Cardinality() (core.Cardinal, bool) {
	c, err := core.CardinalFromInt(s.Len())
	return c, err == nil
}

func (s MapKeySet[M, K, V]) Clone() MapKeySet[M, K, V] {
	return FromMap(maps.Clone(s.Map))
}

func (s MapKeySet[_, Obj, _]) Contains(obj Obj) bool {
	_, exists := s.Map[obj]
	return exists
}

func (s MapKeySet[_, Obj, _]) ForEachUntil(yield func(obj Obj) bool) {
	mapseqs.KeysOf(s.Map).ForEachUntil(yield)
}

func (s MapKeySet[_, _, _]) Len() int {
	return len(s.Map)
}

func (s MapKeySet[_, Obj, _]) Values(yield func(obj Obj) bool) {
	maps.Keys(s.Map)(yield)
}

var _ interface {
	core.CountableSetOf[any]
	seqs.FiniteSeq[any]
} = MapKeySet[map[any]any, any, any]{}
