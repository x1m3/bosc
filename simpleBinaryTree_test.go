package bosc

import (
	"testing"
	"math/rand"
	"fmt"
)

type MyNumber int

func (a MyNumber) Compare(b Comparable) int {
	return int(a.Key().(int) - b.(MyNumber).Key().(int))
}

func (a MyNumber) Key() interface{} {
	return int(a)
}

func TestSimpleBinaryTree(t *testing.T) {

	tree := NewSimpleBinaryTree()
	perms := rand.Perm(1000)
	for _, i := range perms {
		tree.Add(MyNumber(i))
	}

	for i := 0; i < 1000; i++ {
		if _, err := tree.Find(MyNumber(i)); err != nil {
			t.Error("La cagastes Burt Lancaster")
		}
	}

	tree.RangeAll(func(val Comparable) { fmt.Printf("%v ", val) })
	fmt.Println("")

	tree.Range(MyNumber(100), MyNumber(200), func(val Comparable) { fmt.Printf("%v ", val) })
}

func TestSimpleBinaryTree_Remove_NoChilds(t *testing.T) {

	// Let's create a tree unbalanced to the right
	tree := NewSimpleBinaryTree()
	for i := 0; i < 10; i++ {
		tree.Add(MyNumber(i))
	}

	// Removing the last one
	tree.Remove(MyNumber(9))

	values := []MyNumber{}
	tree.RangeAll(func(val Comparable) { values = append(values, val.(MyNumber)) })

	if !testSliceEqual([]MyNumber{0, 1, 2, 3, 4, 5, 6, 7, 8}, values) {
		t.Error("Error removing the a node with no childs")
	}
}

func TestSimpleBinaryTree_Remove_OnlyOneNodeInTree(t *testing.T) {
	tree := NewSimpleBinaryTree()
	tree.Add(MyNumber(5))
	tree.Remove(MyNumber(5))
}

func TestSimpleBinaryTree_Remove_OneChild(t *testing.T) {

	// Let's create an unbalanced tree to the right
	treeRight := NewSimpleBinaryTree()
	for i := 0; i < 10; i++ {
		treeRight.Add(MyNumber(i))
	}

	// Let's create an unbalanced tree to the left
	treeLeft := NewSimpleBinaryTree()
	for i := 9; i >= 0; i-- {
		treeLeft.Add(MyNumber(i))
	}

	// Removing a node in the middle
	treeRight.Remove(MyNumber(5))
	treeLeft.Remove(MyNumber(5))

	// Let's check right tree type
	values := []MyNumber{}
	treeRight.RangeAll(func(val Comparable) { values = append(values, val.(MyNumber)) })
	if !testSliceEqual([]MyNumber{0, 1, 2, 3, 4, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one right child")
	}

	// Let's check left tree type
	values = []MyNumber{}
	treeLeft.RangeAll(func(val Comparable) { values = append(values, val.(MyNumber)) })
	if !testSliceEqual([]MyNumber{0, 1, 2, 3, 4, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one left child")
	}

	// Let's remove the root
	treeLeft.Remove(MyNumber(9))
}

func TestSimpleBinaryTree_Remove2Childs(t *testing.T) {
	tree := NewSimpleBinaryTree()
	tree.Add(MyNumber(5))
	tree.Add(MyNumber(2))
	tree.Add(MyNumber(3))
	tree.Add(MyNumber(4))
	tree.Add(MyNumber(1))
	tree.Add(MyNumber(8))
	tree.Add(MyNumber(7))
	tree.Add(MyNumber(9))
	tree.Add(MyNumber(6))
	tree.Add(MyNumber(0))

	tree.Remove(MyNumber(2))
	values := []MyNumber{}
	tree.RangeAll(func(val Comparable) { values = append(values, val.(MyNumber)) })
	if !testSliceEqual([]MyNumber{0, 1, 3, 4, 5, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one left child")
	}

	// Let's remove the root
	tree.Remove(MyNumber(5))
	values = []MyNumber{}
	tree.RangeAll(func(val Comparable) { values = append(values, val.(MyNumber)) })
	if !testSliceEqual([]MyNumber{0, 1, 3, 4, 6, 7, 8, 9}, values) {
		t.Error("Error removing a node with one left child")
	}
}

func TestSimpleBinaryTree_Range(t *testing.T) {
	ITEMS := 10000000

	tree := NewSimpleBinaryTree()
	for _, j := range rand.Perm(ITEMS) {
		if err := tree.Add(MyNumber(j)); err != nil {
			t.Error(err)
		}
	}

	from := MyNumber(450000)
	to := MyNumber(450020)

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
	fmt.Println(nodes)
}

func TestSimpleBinaryTree_RangeFrom(t *testing.T) {
	ITEMS := 1000000

	tree := NewSimpleBinaryTree()
	for _, j := range rand.Perm(ITEMS) {
		if err := tree.Add(MyNumber(j)); err != nil {
			t.Error(err)
		}
	}

	from := MyNumber(1000)


	tree.RangeFrom(from, func(val Comparable) {
		if val.Compare(from) < 0 {
			t.Errorf("Error in range. Found a value <%d> that shouldn't be in range", val.Key().(int))
		}
	})
}

func TestSimpleBinaryTree_RemoveBrutalRandom(t *testing.T) {

	ITEMS := 100000
	TIMES := 5

	for i := 0; i < TIMES; i++ {

		control := make(map[int]bool, ITEMS)

		// Let's create a tree and fill it with random nodes
		tree := NewSimpleBinaryTree()
		for _, j := range rand.Perm(ITEMS) {
			if err := tree.Add(MyNumber(j)); err != nil {
				t.Error(err)
			}
			control[j] = true
		}

		// Let's check if all nodes are in the tree
		for _, j := range rand.Perm(ITEMS) {
			if _, err := tree.Find(MyNumber(j)); err != nil {
				t.Error(err)
			}
		}

		// Let's remove all nodes randomly
		for _, j := range rand.Perm(ITEMS) {
			if ok := tree.Remove(MyNumber(j)); !ok {
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
func testSliceEqual(slice1 []MyNumber, slice2 []MyNumber) bool {
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
