package core

import "github.com/siliconbrain/go-seqs/seqs"

// SetOf represents a set of objects.
type SetOf[Obj any] interface {
	Contains(obj Obj) bool
}

// CountableSetOf represents a countable set of objects.
// When enumerating a countable set it can yield its objects in any order but must not yield the same object twice.
type CountableSetOf[Obj any] interface {
	SetOf[Obj]
	seqs.Seq[Obj]
}

// Modifiable defines how a set can be modified.
type Modifiable[Obj any] interface {
	// Exclude modifies the collection so that none of the specified objects are present in it.
	Exclude(...Obj)
	// Include modifies the collection so that all of the specified objects are present in it.
	Include(...Obj)
}
