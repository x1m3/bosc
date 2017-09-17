package bosc

import (
	"testing"
	"math/rand"
	"fmt"
)

type myNumber int

func (a myNumber) Compare(b Comparable) int {
	return int(a.Key().(int) - b.(myNumber).Key().(int))
}

func (a myNumber) Key() interface{} {
	return int(a)
}

func TestSimpleBinaryTree(t *testing.T) {

	tree := newSimpleBinaryTree()
	perms := rand.Perm(1000)
	for _, i := range perms {
		tree.Add(myNumber(i))
	}

	for i := 0; i < 1000; i++ {
		if _, err := tree.Find(myNumber(i)); err != nil {
			t.Error("La cagastes Burt Lancaster")
		}
	}

	tree.RangeAll(func(val Comparable) { fmt.Printf("%v ", val) })
	fmt.Println("")

	tree.Range(myNumber(100), myNumber(200), func(val Comparable) { fmt.Printf("%v ", val) })
}

func TestSimpleBinaryTree_Remove_NoChilds(t *testing.T) {

	// Let's create a tree unbalanced to the right
	tree := newSimpleBinaryTree()
	for i := 0; i < 10; i++ {
		tree.Add(myNumber(i))
	}

	// Removing the last one
	tree.Remove(myNumber(9))

	values := []myNumber{}
	tree.RangeAll(func(val Comparable) { values = append(values, val.(myNumber)) })

	if !testSliceEqual([]myNumber{0, 1, 2, 3, 4, 5, 6, 7, 8}, values) {
		t.Error("Error removing the a node with no childs")
	}
}

func TestSimpleBinaryTree_Remove_OnlyOneNodeInTree(t *testing.T) {
	tree := newSimpleBinaryTree()
	tree.Add(myNumber(5))
	tree.Remove(myNumber(5))
}

func TestSimpleBinaryTree_Remove_OneChild(t *testing.T) {

	// Let's create an unbalanced tree to the right
	treeRight := newSimpleBinaryTree()
	for i := 0; i < 10; i++ {
		treeRight.Add(myNumber(i))
	}

	// Let's create an unbalanced tree to the left
	treeLeft := newSimpleBinaryTree()
	for i := 9; i >= 0; i-- {
		treeLeft.Add(myNumber(i))
	}

	// Removing a node in the middle
	treeRight.Remove(myNumber(5))
	treeLeft.Remove(myNumber(5))

	// Let's check right tree type
	values := []myNumber{}
	treeRight.RangeAll(func(val Comparable) { values = append(values, val.(myNumber)) })
	if !testSliceEqual([]myNumber{0, 1, 2, 3, 4, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one right child")
	}

	// Let's check left tree type
	values = []myNumber{}
	treeLeft.RangeAll(func(val Comparable) { values = append(values, val.(myNumber)) })
	if !testSliceEqual([]myNumber{0, 1, 2, 3, 4, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one left child")
	}

	// Let's remove the root
	treeLeft.Remove(myNumber(9))
}

func TestSimpleBinaryTree_Remove2Childs(t *testing.T) {
	tree := newSimpleBinaryTree()
	tree.Add(myNumber(5))
	tree.Add(myNumber(2))
	tree.Add(myNumber(3))
	tree.Add(myNumber(4))
	tree.Add(myNumber(1))
	tree.Add(myNumber(8))
	tree.Add(myNumber(7))
	tree.Add(myNumber(9))
	tree.Add(myNumber(6))
	tree.Add(myNumber(0))

	tree.Remove(myNumber(2))
	values := []myNumber{}
	tree.RangeAll(func(val Comparable) { values = append(values, val.(myNumber)) })
	if !testSliceEqual([]myNumber{0, 1, 3, 4, 5, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one left child")
	}

	// Let's remove the root
	tree.Remove(myNumber(5))
	values = []myNumber{}
	tree.RangeAll(func(val Comparable) { values = append(values, val.(myNumber)) })
	if !testSliceEqual([]myNumber{0, 1, 3, 4, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one left child")
	}
}

func TestSimpleBinaryTree_Range(t *testing.T) {
	ITEMS := 1000000

	tree := newSimpleBinaryTree()
	for _, j := range rand.Perm(ITEMS) {
		if err := tree.Add(myNumber(j)); err != nil {
			t.Error(err)
		}
	}

	from := myNumber(45000)
	to := myNumber(46000)

	nodes := 0
	tree.Range(from, to, func(val Comparable) {
		nodes++
		if val.Compare(from) < 0 {
			t.Errorf("Error in range. Found a value <%d> that shouldn't be in range", val.Key().(int))
		}
		if val.Compare(to) > 0 {
			t.Errorf("Error in range. Found a value <%d> that shouldn't be in range", val.Key().(int))
		}
	})
	nodes_expected :=  to.Key().(int) - from.Key().(int) + 1
	if nodes != nodes_expected {
		t.Errorf("Range function executed over %d nodes. Expecting %d", nodes, nodes_expected)
	}
}

func TestSimpleBinaryTree_RangeFrom(t *testing.T) {
	ITEMS := 1000000

	tree := newSimpleBinaryTree()
	for _, j := range rand.Perm(ITEMS) {
		if err := tree.Add(myNumber(j)); err != nil {
			t.Error(err)
		}
	}

	from := myNumber(1000)

	nodes := 0
	tree.RangeFrom(from, func(val Comparable) {
		nodes++
		if val.Compare(from) < 0 {
			t.Errorf("Error in range. Found a value <%d> that shouldn't be in range", val.Key().(int))
		}
	})
	nodes_expected :=  tree.Max().Key().(int) - from.Key().(int) + 1
	if nodes != nodes_expected {
		t.Errorf("Range function executed over %d nodes. Expecting %d", nodes, nodes_expected)
	}
}

func TestSimpleBinaryTree_RangeTo(t *testing.T) {
	ITEMS := 1000000

	tree := newSimpleBinaryTree()
	for _, j := range rand.Perm(ITEMS) {
		if err := tree.Add(myNumber(j)); err != nil {
			t.Error(err)
		}
	}

	to := myNumber(5000)

	nodes := 0
	tree.RangeTo(to, func(val Comparable) {
		nodes++
		if val.Compare(to) > 0 {
			t.Errorf("Error in range. Found a value <%d> that shouldn't be in range", val.Key().(int))
		}
	})
	nodes_expected :=  to.Key().(int)  - tree.Min().Key().(int) + 1
	if nodes != nodes_expected {
		t.Errorf("Range function executed over %d nodes. Expecting %d", nodes, nodes_expected)
	}
}

func TestSimpleBinaryTree_RemoveBrutalRandom(t *testing.T) {

	ITEMS := 100000
	TIMES := 5

	for i := 0; i < TIMES; i++ {

		control := make(map[int]bool, ITEMS)

		// Let's create a tree and fill it with random nodes
		tree := newSimpleBinaryTree()
		for _, j := range rand.Perm(ITEMS) {
			if err := tree.Add(myNumber(j)); err != nil {
				t.Error(err)
			}
			control[j] = true
		}

		// Let's check if all nodes are in the tree
		for _, j := range rand.Perm(ITEMS) {
			if _, err := tree.Find(myNumber(j)); err != nil {
				t.Error(err)
			}
		}

		// Let's remove all nodes randomly
		for _, j := range rand.Perm(ITEMS) {
			if ok := tree.Remove(myNumber(j)); !ok {
				t.Errorf("Unkown error removing item <%v>, loop:<%v>", j, i)
			}
			delete(control, j)

			// Let's do some checks from time to time
			if j%10000 == 0 || j < 100 {
				// Tree has correct len?
				if tree.Count() != uint64(len(control)) {
					t.Errorf("Wrong tree lenght. Error removing. Expecting %d items, got %d", len(control), tree.Count())
				}

				// All items in tree are inside control?
				tree.RangeAll(func(val Comparable) {
					if _, ok := control[val.Key().(int)]; !ok {
						t.Errorf("Error removing items. %d removed unnintentionally", val.Key().(int))
					}
				})
			}
		}
	}
}

// Returns true if 2 slices are equal, in length, values and order.
// false otherwise
func testSliceEqual(slice1 []myNumber, slice2 []myNumber) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range (slice1) {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
