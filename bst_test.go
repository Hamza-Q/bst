// +build go1.18

package bst

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

// TestInsert tests BST.Insert.
func TestInsert(t *testing.T) {
	tree := &BST[int]{}
	root := &node[int]{value: 5}
	tree.root = root
	t.Cleanup(func() {
		if t.Failed() {
			t.Log("Final structure of tree:")
			tree.Inorder(func(val int) bool{
				fmt.Print(val, "")
				return false
			})
		}
	})
	tree.Insert(6)
	if root.right == nil {
		t.Fatal("root.right is nil")
	}
	if root.right.value != 6 {
		t.Fatalf("root.right.value expected to be 6, got %d", root.right.value)
	}
	tree.Insert(7)
	if root.right.right == nil {
		t.Fatalf("Bad insertion")
	}
}


// getSorted returns a tree-sorted slice of the elements in the subtree of root.
func getSorted[T Ordered](root *node[T]) []T {
	sorted := []T{}
	root.inorder(func(val T) bool {
		sorted = append(sorted, val)
		return false
	})
	return sorted
}

// TestInsertSorted tests that a BST can be used to sort a permuted slice of
// ints.
func TestInsertSorted(t *testing.T) {
	rand.Seed(42)
	vals := rand.Perm(1000)
	sortedVals := make([]int, len(vals))
	copy(sortedVals, vals)
	sort.Ints(sortedVals)

	root := &node[int]{value: vals[0]}
	for _, val := range vals[1:] {
		insert(root, val)
	}

	treeSorted := getSorted(root)
	if !reflect.DeepEqual(sortedVals, treeSorted) {
		t.Fatalf("Sorted and inorder not equal.\nGot: %v\n", treeSorted)
	}
}

// TestInsertEmpty tests insertion when the root is empty.
func TestInsertEmpty(t *testing.T) {
	tree := &BST[int]{}
	tree.Insert(10)
	if tree.root == nil {
		t.Fatalf("Expected tree to have a root; Tree:\n %+v", tree)
	}
	if !tree.Get(10) {
		t.Fatal("Could not find inserted node")
	}
}


// TestInorderStopIter tests that inorder returns early successfully.
func TestInorderStopIter(t *testing.T) {
	tree := &BST[int]{}
	tree.root = &node[int]{
		value: 20,
	}
	for _, val := range []int{5, 10, 7, 13, 17} {
		tree.Insert(val)
	}
	for _, val := range []int{25, 23, 27, 35, 30} {
		tree.Insert(val)
	}
	visited := []int{}
	visitFunc := func(val int) bool {
		visited = append(visited, val)
		if val > 20 {
			return true
		}
		return false
	}
	tree.Inorder(visitFunc)
	// Expect to visit the first value greater than 20 before stopping iteration.
	expectedVisit := []int{5, 7, 10, 13, 17, 20, 23}
	if !reflect.DeepEqual(visited, expectedVisit) {
		t.Fatalf("Expected to visit %v\nActually visited %v", expectedVisit, visited)
	}
}
