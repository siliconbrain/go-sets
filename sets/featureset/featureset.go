package featureset

import (
	"github.com/siliconbrain/go-sets/sets/core"
)

func From[Set core.SetOf[Feat], Obj, Feat any](set Set, getFeature func(Obj) Feat) FeatureSet[Set, Obj, Feat] {
	if getFeature == nil {
		panic("feature extraction function must not be nil")
	}
	return FeatureSet[Set, Obj, Feat]{
		BaseSet:    set,
		GetFeature: getFeature,
	}
}

func FromModifiable[
	Set interface {
		core.SetOf[Feat]
		core.Modifiable[Feat]
	},
	Obj, Feat any,
](set Set, getFeature func(Obj) Feat) ModifiableFeatureSet[Set, Obj, Feat] {
	return ModifiableFeatureSet[Set, Obj, Feat]{
		FeatureSet: From(set, getFeature),
	}
}

type FeatureSet[Set core.SetOf[Feat], Obj, Feat any] struct {
	BaseSet    Set
	GetFeature func(Obj) Feat
}

func (s FeatureSet[_, Obj, _]) Contains(obj Obj) bool {
	return s.BaseSet.Contains(s.GetFeature(obj))
}

type ModifiableFeatureSet[
	Set interface {
		core.SetOf[Feat]
		core.Modifiable[Feat]
	},
	Obj, Feat any,
] struct {
	FeatureSet[Set, Obj, Feat]
}

func (s ModifiableFeatureSet[_, Obj, _]) Exclude(objs ...Obj) {
	s.BaseSet.Exclude(s.getFeatures(objs)...)
}

func (s ModifiableFeatureSet[_, Obj, _]) Include(objs ...Obj) {
	s.BaseSet.Include(s.getFeatures(objs)...)
}

func (s ModifiableFeatureSet[_, Obj, Feat]) getFeatures(objs []Obj) (feats []Feat) {
	if l := len(objs); l > 0 {
		feats = make([]Feat, l)
		for i := range objs {
			feats[i] = s.GetFeature(objs[i])
		}
	}
	return
}
