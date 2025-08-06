package funcset

import "github.com/siliconbrain/go-sets/sets/core"

// FromFunc returns a FuncSet implemented by the specified function.
func FromFunc[Obj any](contains func(Obj) bool) FuncSet[Obj] {
	return contains
}

// FuncSet implements [core.SetOf] using a function.
type FuncSet[Obj any] func(Obj) bool

// Contains returns true when the underlying function returns true for the specified object.
func (set FuncSet[Obj]) Contains(obj Obj) bool {
	return set(obj)
}

var _ core.SetOf[any] = FuncSet[any](nil)
