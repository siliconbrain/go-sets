package arrayset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySet_Contains(t *testing.T) {
	testCases := map[string]struct {
		items []int
		value int
		want  bool
	}{
		"empty set": {
			items: []int{},
			value: 42,
			want:  false,
		},
		"internal": {
			items: []int{1, 2, 3, 4},
			value: 2,
			want:  true,
		},
		"external": {
			items: []int{1, 2, 3, 4},
			value: 0,
			want:  false,
		},
		"duplicate items, internal": {
			items: []int{1, 2, 3, 4, 1, 2, 3, 4},
			value: 3,
			want:  true,
		},
		"duplicate items, external": {
			items: []int{1, 2, 3, 4, 1, 2, 3, 4},
			value: 0,
			want:  false,
		},
	}
	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			set := FromSlice(testCase.items)
			assert.Equal(t, testCase.want, set.Contains(testCase.value))
		})
	}
}

func TestArraySet_Exclude(t *testing.T) {
	testCases := map[string]struct {
		items     []int
		exclude   []int
		wantItems []int
	}{
		"empty set": {
			items:     []int{},
			exclude:   []int{42},
			wantItems: []int{},
		},
		"orthogonal sets": {
			items:     []int{1, 2, 3, 4},
			exclude:   []int{42, 21},
			wantItems: []int{1, 2, 3, 4},
		},
		"exclude partial intersection": {
			items:     []int{1, 2, 3, 4},
			exclude:   []int{0, 2, 4, 6},
			wantItems: []int{1, 3},
		},
		"exclude superset": {
			items:     []int{1, 2, 3, 4},
			exclude:   []int{0, 1, 2, 3, 4, 5},
			wantItems: []int{},
		},
		"exclude subset": {
			items:     []int{1, 2, 3, 4},
			exclude:   []int{1, 4},
			wantItems: []int{2, 3},
		},
		"exclude empty set": {
			items:     []int{1, 2, 3, 4},
			exclude:   []int{},
			wantItems: []int{1, 2, 3, 4},
		},
	}
	for name, testCase := range testCases {
		testCase := testCase
		t.Run(name, func(t *testing.T) {
			set := FromSlice(testCase.items)
			set.Exclude(testCase.exclude...)
			assert.ElementsMatch(t, testCase.wantItems, set.Items)
		})
	}
}
