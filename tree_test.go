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
