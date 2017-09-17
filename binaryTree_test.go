package bosc_test

import (
	"testing"
	"bosc"
	"fmt"
	"math/rand"
)

type myIndex struct {
	key int
	value string
}

func (a myIndex) Compare(b bosc.Comparable) int {
	return a.Key().(int) - b.(bosc.Comparable).Key().(int)
}

func (a myIndex) Key() interface{} {
	return a.key
}

func newIndex(n int) *myIndex {
	return &myIndex{key:n, value:fmt.Sprintf("Number %d", n)}
}

func Test_Add_Find(t *testing.T) {
	for _, treeType := range bosc.TreeTypes() {
		tree, err := bosc.NewTree(treeType)
		if err!=nil {
			t.Errorf("Tree type <%s>. Error creating tree <%v>", treeType, err)
		}

		// Add values
		for i:= range rand.Perm(5000) {
			tree.Add(newIndex(i))
		}

		// Search values
		for i:= range rand.Perm(5000) {
			if index, err := tree.Find(&myIndex{key:i}); err!=nil {
				t.Errorf("Tree type <%s>. <%s>",treeType, err)
			}else {
				expected := fmt.Sprintf("Number %d", i)
				if index.(*myIndex).value != expected {
					t.Errorf("Tree type <%s>, Error retriving key. Expecting a value of <%s>, got <%s>", treeType, expected, index.(*myIndex).value)
				}
			}
		}
	}
}
