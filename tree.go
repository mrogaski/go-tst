// Package tst implements a generic ternary search tree.
package tst

// ByteSequence is a constraint interface for the search tree key.
type ByteSequence interface {
	~string | ~[]byte
}

// TernarySearchTree is a structure implementing Sedgewick and Bentley's Ternary Search Tree.
type TernarySearchTree[K ByteSequence, V any] struct{}

// New returns an initialized, empty search tree.
func New[K ByteSequence, V any]() *TernarySearchTree[K, V] {
	return new(TernarySearchTree[K, V])
}
