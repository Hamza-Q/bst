// +build go1.18

package bst

// Ordered is a type constraint that matches all ordered types.
// (An ordered type is one that supports the < <= >= > operators.)
// To be replaced by a more canonical interface.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// BST is an implementation of a binary search tree.
type BST[T Ordered] struct {
	size int
	root *node[T]
}

// Insert adds the given key to the tree. 
func (b *BST[T]) Insert(value T) {
	b.size++
	if b.root == nil {
		b.root = &node[T]{
			value: value,
		}
		return
	}
	insert(b.root, value)
}

// Inorder performs an inorder traversal of the binary search tree.
// The visit callback is performed at every node unless it returns true,
// which stops iteration.
func (b *BST[T]) Inorder(visit func(val T) bool) {
	if b.root == nil {
		return
	}
	b.root.inorder(visit)
}

// Get returns whether or not the given value exists in the tree.
func (b *BST[T]) Get(value T) (ok bool) {
	if b.root == nil {
		return false
	}
	for n := b.root; n != nil; {
		if value == n.value {
			return true
		}
		if value > n.value {
			n = n.right
		} else {
			n = n.left
		}
	}
	return false
}

// node is a node of a Binary Search Tree.
//
// TODO: Move beyond the Ordered generic and support custom types that implement
// the "sort" package interfaces.
// Reasonable options are to duplicate/generate this code for a constraint
// allowing a "Lesser" interface type, or to include a Lesser interface type and
// do some type switching (see https://github.com/golang/go/issues/45380)
type node[T Ordered] struct {
	parent *node[T]
	left   *node[T]
	right  *node[T]
	value  T
}

// insert creates a node with the given value, inserted where it belongs in the
// tree.
func insert[T Ordered](root *node[T], value T) *node[T] {
	// TODO: Should this structure support multiple keys of the same value?
	// Should it be configurable?
	if value >= root.value {
		if root.right != nil {
			return insert(root.right, value)
		}
		root.right = &node[T]{
			parent: root,
			value:  value,
		}
		return root.right
	}
	if root.left != nil {
		return insert(root.left, value)
	}
	root.left = &node[T]{
		parent: root,
		value:  value,
	}
	return root.left
}

// inorder performs an inorder dfs traversal of the subtree rooted at this node,
// stopping iteration if the given visit function returns true.
// The visit function is called at each visited node.
func (n *node[T]) inorder(visit func(value T) bool) (stopIter bool) {
	if n == nil {
		return false
	}
	if n.left.inorder(visit) {
		return true
	}
	if visit(n.value) {
		return true
	}
	return n.right.inorder(visit)
}
