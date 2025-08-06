package core

import "github.com/siliconbrain/go-seqs/seqs"

// QuickCardinalityOf returns the cardinality of the set if it can be determined without counting its elements.
func QuickCardinalityOf[Obj any](set SetOf[Obj]) (Cardinal, bool) {
	if s, ok := set.(interface{ Cardinality() (Cardinal, bool) }); ok {
		return s.Cardinality()
	}
	if lener, ok := set.(seqs.Lener); ok {
		c, err := CardinalFromInt(lener.Len())
		return c, err == nil
	}
	return 0, false
}
