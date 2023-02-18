package tst_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrogaski/go-tst"
)

func TestNew(t *testing.T) {
	t.Parallel()

	got := tst.New[string, int]()

	assert.IsType(t, &tst.TernarySearchTree[string, int]{}, got)
	assert.NotNil(t, got)
}

func TestTernarySearchTree_Insert(t *testing.T) {
	t.Parallel()

	type element struct {
		k string
		v int
	}

	tests := []struct {
		name    string
		preload []element
		key     string
		value   int
		want    bool
	}{
		{
			name:  "initital",
			key:   "firewater",
			value: 1,
			want:  false,
		},
		{
			name: "new",
			preload: []element{
				{k: "fire", v: 1},
			},
			key:   "firewater",
			value: 2,
			want:  false,
		},
		{
			name: "update",
			preload: []element{
				{k: "firewater", v: 1},
			},
			key:   "firewater",
			value: 2,
			want:  true,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tree := tst.New[string, int]()

			if tc.preload != nil {
				for _, elem := range tc.preload {
					tree.Insert(elem.k, elem.v)
				}
			}

			got := tree.Insert(tc.key, tc.value)

			assert.Equal(t, tc.want, got)
		})
	}
}
