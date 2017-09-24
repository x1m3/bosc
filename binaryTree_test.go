package bosc_test

import (
	"bosc"
	"fmt"
	"math/rand"
	"testing"
)

type myIndex struct {
	key   int
	value string
}

func (a myIndex) Compare(b bosc.Comparable) int {
	return a.Key().(int) - b.(bosc.Comparable).Key().(int)
}

func (a myIndex) Key() interface{} {
	return a.key
}

func newIndex(n int) *myIndex {
	return &myIndex{key: n, value: fmt.Sprintf("Number %d", n)}
}

func Test_Add_Find(t *testing.T) {
	for _, treeType := range bosc.TreeTypes() {
		tree, err := bosc.NewTree(treeType)
		if err != nil {
			t.Errorf("Tree type <%s>. Error creating tree <%v>", treeType, err)
		}

		// Add values
		for _, i := range rand.Perm(5000) {
			tree.Add(newIndex(i))
		}

		// Search values
		for _, i := range rand.Perm(5000) {
			if index, err := tree.Find(&myIndex{key: i}); err != nil {
				t.Errorf("Tree type <%s>. <%s>", treeType, err)
			} else {
				expected := fmt.Sprintf("Number %d", i)
				if index.(*myIndex).value != expected {
					t.Errorf("Tree type <%s>, Error retriving key. Expecting a value of <%s>, got <%s>", treeType, expected, index.(*myIndex).value)
				}
			}
		}
	}
}

func Test_Min_Max(t *testing.T) {
	for _, treeType := range bosc.TreeTypes() {
		tree, err := bosc.NewTree(treeType)
		if err != nil {
			t.Errorf("Tree type <%s>. Error creating tree <%v>", treeType, err)
		}

		// Add values
		for _, i := range rand.Perm(5000) {
			tree.Add(newIndex(i))
		}
		min := tree.Min().Key().(int)
		max := tree.Max().Key().(int)

		if min != 0 {
			t.Errorf("Tree type <%s>. Wrong min value. Expecting <0>, got <%d>", treeType, min)
		}
		if max != 4999 {
			t.Errorf("Tree type <%s>. Wrong max value. Expecting <5000>, got <%d>", treeType, max)
		}
	}

}

func Test_Add_Remove(t *testing.T) {
	for _, treeType := range bosc.TreeTypes() {
		tree, err := bosc.NewTree(treeType)
		if err != nil {
			t.Errorf("Tree type <%s>. Error creating tree <%v>", treeType, err)
		}

		// Add values
		for _, i := range rand.Perm(5000) {
			tree.Add(newIndex(i))
		}

		// Removing values
		for _, i := range rand.Perm(5000) {
			if !tree.Remove(newIndex(i)) {
				t.Errorf("Tree type <%s>. Error removing item <%d>. Item not found", treeType, i)
			}

			if _, err := tree.Find(newIndex(i)); err == nil {
				t.Errorf("Tree type <%s>. Error removing item <%d>: <%s>", treeType, newIndex(i), err)
			}
		}
	}
}

func Test_Ranges(t *testing.T) {
	for _, treeType := range bosc.TreeTypes() {
		tree, err := bosc.NewTree(treeType)
		if err != nil {
			t.Errorf("Tree type <%s>. Error creating tree <%v>", treeType, err)
		}

		// Add values
		for _, i := range rand.Perm(5000) {
			tree.Add(newIndex(i))
		}

		// Removing values
		last := -1
		tree.RangeAll(func(c bosc.Comparable) {
			if c.Key().(int) < 0 || c.Key().(int) > 5000 {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll <%d>", treeType, c.Key().(int))
			}
			if c.Key().(int) <= last {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll. Expecting <%d>, got <%d>", treeType, last, c.Key().(int))
			}
			last = c.Key().(int)
		})

		last = 2499
		tree.RangeFrom(newIndex(2500), func(c bosc.Comparable) {
			if c.Key().(int) < 2500 || c.Key().(int) > 5000 {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll <%d>", treeType, c.Key().(int))
			}
			if c.Key().(int) <= last {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll. Expecting <%d>, got <%d>", treeType, last, c.Key().(int))
			}
			last = c.Key().(int)
		})

		last = -1
		tree.RangeTo(newIndex(2500), func(c bosc.Comparable) {
			if c.Key().(int) < 0 || c.Key().(int) > 2500 {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll <%d>", treeType, c.Key().(int))
			}
			if c.Key().(int) <= last {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll. Expecting <%d>, got <%d>", treeType, last, c.Key().(int))
			}
			last = c.Key().(int)
		})

		last = 999
		tree.Range(newIndex(1000), newIndex(4000), func(c bosc.Comparable) {
			if c.Key().(int) < 1000 || c.Key().(int) > 4000 {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll <%d>", treeType, c.Key().(int))
			}
			if c.Key().(int) <= last {
				t.Errorf("Tree type <%s>. Found a wrong value in RangeAll. Expecting <%d>, got <%d>", treeType, last, c.Key().(int))
			}
			last = c.Key().(int)
		})
	}
}
