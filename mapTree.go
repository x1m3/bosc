package bosc

import "errors"

func init() {
	register("map_tree", func() BinarySearchTree { return newMapTree() })
}

type mapTree struct {
	m map[interface{}]Comparable
}

func newMapTree() *mapTree {
	t := new(mapTree)
	t.m = make(map[interface{}]Comparable, 1024)
	return t
}

func (t *mapTree) Add(val Comparable) error {
	if _,exists := t.m[val.Key()];exists {
		return errors.New("Duplicated value.")
	}
	t.m[val.Key()]=val
	return nil
}

func (t *mapTree) Count() uint64 {
	return uint64(len(t.m))
}

func (t *mapTree) Find(val Comparable) (Comparable, error) {
	if node,exists := t.m[val.Key()];exists {
		return node, nil
	}
	return nil, errors.New("Not Found.")
}


func (t *mapTree) Min() Comparable {return nil}
func (t *mapTree) Max() Comparable{return nil}
func (t *mapTree) Remove(val Comparable) (found bool) { return false}
func (t *mapTree) Range(valFrom Comparable, valTo Comparable, fn func(node Comparable)) {}
func (t *mapTree) RangeAll(fn func(node Comparable)) {}
func (t *mapTree) RangeFrom(val Comparable, fn func(val Comparable)) {}
func (t *mapTree) RangeTo(val Comparable, fn func(val Comparable)) {}
