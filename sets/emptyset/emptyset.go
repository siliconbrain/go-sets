package emptyset

import "github.com/siliconbrain/go-sets/sets/core"

type EmptySet[Obj any] struct{}

func (EmptySet[Obj]) Cardinality() (core.Cardinal, bool) {
	return 0, true
}

func (EmptySet[Obj]) Contains(Obj) bool {
	return false
}

func (EmptySet[Obj]) ForEachUntil(func(Obj) bool) {}

func (EmptySet[Obj]) Len() int {
	return 0
}

func (EmptySet[Obj]) Values(func(Obj) bool) {}
