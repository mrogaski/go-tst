// Package tst implements a generic ternary search tree.
package tst

// ByteSequence is a constraint interface for the search tree key.
type ByteSequence interface {
	~string | ~[]byte
}

// TernarySearchTree is a structure implementing Sedgewick and Bentley's Ternary Search Tree.
type TernarySearchTree[K ByteSequence, V any] struct {
	root *node[V]
}

type node[V any] struct {
	splitChar byte
	loKid     *node[V]
	eqKid     *node[V]
	hiKid     *node[V]
}

// New returns an initialized, empty search tree.
// K is a string or byte slice type for the keys and V is any type for the values.
func New[K ByteSequence, V any]() *TernarySearchTree[K, V] {
	return new(TernarySearchTree[K, V])
}

// Insert adds a value to the tree, returning true if the value is updating a previous value or false if this
// is the first value inserted with the key.
func (t *TernarySearchTree[K, V]) Insert(key K, value V) bool {
	sequence := []byte(key)

	nodePtr := &t.root

	updated := true

	i := 0

	for {
		if i >= len(sequence) {
			break
		}

		char := sequence[i]

		if *nodePtr == nil {
			*nodePtr = &node[V]{splitChar: char}
			updated = false
		}

		switch {
		case char < (*nodePtr).splitChar:
			nodePtr = &(*nodePtr).loKid
		case char > (*nodePtr).splitChar:
			nodePtr = &(*nodePtr).hiKid
		default:
			nodePtr = &(*nodePtr).eqKid
			i++
		}
	}

	return updated
}
