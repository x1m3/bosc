package bosc

import (
	"errors"
	"sort"
)

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
	if _, exists := t.m[val.Key()]; exists {
		return errors.New("Duplicated value.")
	}
	t.m[val.Key()] = val
	return nil
}

func (t *mapTree) Count() uint64 {
	return uint64(len(t.m))
}

func (t *mapTree) Find(val Comparable) (Comparable, error) {
	if node, exists := t.m[val.Key()]; exists {
		return node, nil
	}
	return nil, errors.New("Not Found.")
}

func (t *mapTree) Min() Comparable {
	var min Comparable

	for _, value := range t.m {
		if min != nil {
			if min.Compare(value) >= 0 {
				min = value
			}
		} else {
			min = value
		}
	}
	return min
}

func (t *mapTree) Max() Comparable {
	var max Comparable

	for _, value := range t.m {
		if max != nil {
			if max.Compare(value) <= 0 {
				max = value
			}
		} else {
			max = value
		}
	}
	return max
}

func (t *mapTree) Remove(val Comparable) (found bool) {
	if _, found := t.m[val.Key()]; found {
		delete(t.m, val.Key())
		return true
	}
	return false
}

func (t *mapTree) Range(valFrom Comparable, valTo Comparable, fn func(node Comparable)) {
	list := t.newOrderedList()
	for _, i := range list.items {
		if valFrom.Compare(i) >= 0 && valTo.Compare(i) <= 0 {
			fn(i)
		}
	}
}

func (t *mapTree) RangeAll(fn func(node Comparable)) {
	list := t.newOrderedList()
	for _, i := range list.items {
		fn(i)
	}
}

func (t *mapTree) RangeFrom(val Comparable, fn func(val Comparable)) {
	list := t.newOrderedList()
	for _, i := range list.items {
		if i.Compare(val) >= 0 {
			fn(i)
		}
	}
}

func (t *mapTree) RangeTo(val Comparable, fn func(val Comparable)) {
	list := t.newOrderedList()
	for _, i := range list.items {
		if i.Compare(val) <= 0 {
			fn(i)
		}
	}
}

func (t *mapTree) newOrderedList() *orderedList {
	l := new(orderedList)
	l.items = make([]Comparable, 0, t.Count())
	for _, k := range t.m {
		l.items = append(l.items, k)
	}
	sort.Sort(l)
	return l
}

type orderedList struct {
	items []Comparable
}

func (l *orderedList) Len() int { return len(l.items) }
func (l *orderedList) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
}
func (l *orderedList) Less(i, j int) bool {
	p1 := l.items[i]
	p2 := l.items[j]
	if p1.Compare(p2) < 0 {
		return true
	} else {
		return false
	}
}
