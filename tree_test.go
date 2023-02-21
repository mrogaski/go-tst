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
			key:   "foo",
			value: 1,
			want:  false,
		},
		{
			name: "new",
			preload: []element{
				{k: "foo", v: 1},
			},
			key:   "foobar",
			value: 2,
			want:  false,
		},
		{
			name: "update",
			preload: []element{
				{k: "foobar", v: 1},
			},
			key:   "foobar",
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

func TestTernarySearchTree_Get(t *testing.T) {
	t.Parallel()

	type element struct {
		k string
		v int
	}

	tests := []struct {
		name      string
		preload   []element
		key       string
		want      int
		wantFound bool
	}{
		{
			name:      "empty",
			key:       "foo",
			wantFound: false,
		},
		{
			name: "hit first",
			preload: []element{
				{k: "foo", v: 1},
				{k: "bar", v: 2},
				{k: "baz", v: 3},
			},
			key:       "foo",
			want:      1,
			wantFound: true,
		},
		{
			name: "hit middle",
			preload: []element{
				{k: "foo", v: 1},
				{k: "bar", v: 2},
				{k: "baz", v: 3},
			},
			key:       "bar",
			want:      2,
			wantFound: true,
		},
		{
			name: "hit last",
			preload: []element{
				{k: "foo", v: 1},
				{k: "bar", v: 2},
				{k: "baz", v: 3},
			},
			key:       "baz",
			want:      3,
			wantFound: true,
		},
		{
			name: "miss",
			preload: []element{
				{k: "foo", v: 1},
				{k: "bar", v: 2},
				{k: "baz", v: 3},
			},
			key:       "quux",
			wantFound: false,
		},
		{
			name: "miss substring",
			preload: []element{
				{k: "foobar", v: 1},
				{k: "bar", v: 2},
				{k: "baz", v: 3},
			},
			key:       "foo",
			wantFound: false,
		},
		{
			name: "miss superstring",
			preload: []element{
				{k: "foo", v: 1},
				{k: "bar", v: 2},
				{k: "baz", v: 3},
			},
			key:       "foobar",
			wantFound: false,
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

			got, found := tree.Get(tc.key)

			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.wantFound, found)
		})
	}
}
