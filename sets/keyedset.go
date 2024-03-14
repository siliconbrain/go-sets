package sets

import (
	"github.com/siliconbrain/go-mapseqs/mapseqs"
	"github.com/siliconbrain/go-seqs/seqs"
)

func NewKeyedSet[K comparable, V any](keyFn func(V) K) KeyedSet[K, V] {
	if keyFn == nil {
		panic("keyFn must not be nil")
	}
	return KeyedSet[K, V]{
		keyFn: keyFn,
	}
}

func KeyedSetFromSeq[K comparable, V any](seq seqs.Seq[V], keyFn func(V) K) KeyedSet[K, V] {
	s := NewKeyedSet(keyFn)
	s.IncludeSeq(seq)
	return s
}

func KeyedSetFromSlice[K comparable, V any](vs []V, keyFn func(V) K) KeyedSet[K, V] {
	s := NewKeyedSet(keyFn)
	s.Include(vs...)
	return s
}

func KeyedSetFromValues[K comparable, V any](keyFn func(V) K, vs ...V) (res KeyedSet[K, V]) {
	return KeyedSetFromSlice(vs, keyFn)
}

type KeyedSet[K comparable, V any] struct {
	hashmap map[K]V
	keyFn   func(V) K
}

func (s KeyedSet[K, V]) Cardinality() int {
	return len(s.hashmap)
}

func (s KeyedSet[K, V]) Contains(v V) bool {
	_, res := s.hashmap[s.keyFn(v)]
	return res
}

func (s *KeyedSet[K, V]) Exclude(vs ...V) {
	if len(s.hashmap) == 0 {
		return
	}
	for _, v := range vs {
		delete(s.hashmap, s.keyFn(v))
		if len(s.hashmap) == 0 {
			return
		}
	}
}

func (s *KeyedSet[K, V]) ExcludeSeq(seq seqs.Seq[V]) {
	if len(s.hashmap) == 0 {
		return
	}
	seq.ForEachUntil(func(v V) bool {
		delete(s.hashmap, s.keyFn(v))
		return len(s.hashmap) == 0
	})
}

func (s KeyedSet[K, V]) ForEachUntil(fn func(V) bool) {
	mapseqs.ValuesOf(s.hashmap).ForEachUntil(fn)
}

func (s KeyedSet[K, V]) GetByKey(key K) (V, bool) {
	val, set := s.hashmap[key]
	return val, set
}

func (s *KeyedSet[K, V]) Include(vs ...V) {
	s.ensureHashmap()

	for _, v := range vs {
		s.hashmap[s.keyFn(v)] = v
	}
}

func (s *KeyedSet[K, V]) IncludeSeq(seq seqs.Seq[V]) {
	s.ensureHashmap()

	seqs.ForEach(seq, func(v V) {
		s.hashmap[s.keyFn(v)] = v
	})
}

func (s *KeyedSet[K, V]) ensureHashmap() {
	if s.hashmap == nil {
		s.hashmap = make(map[K]V)
	}
}

var _ interface {
	CountableSetOf[any]
	Modifiable[any]
} = (*KeyedSet[any, any])(nil)
