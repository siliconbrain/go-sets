package sets

import "github.com/siliconbrain/go-sets/sets/core"

type CountableSetOf[Obj any] = core.CountableSetOf[Obj]
type SetOf[Obj any] = core.SetOf[Obj]
type Modifiable[Obj any] = core.Modifiable[Obj]
type Cardinal = core.Cardinal

var CardinalFromInt = core.CardinalFromInt

func QuickCardinalityOf[Obj any](set SetOf[Obj]) (Cardinal, bool) {
	return core.QuickCardinalityOf(set)
}
