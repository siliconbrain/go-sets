package universalset

import "github.com/siliconbrain/go-sets/sets/core"

type UniversalSet[Obj any] struct{}

func (UniversalSet[Obj]) Cardinality() (core.Cardinal, bool) {
	return 0, false
}

func (UniversalSet[Obj]) Contains(Obj) bool {
	return true
}
