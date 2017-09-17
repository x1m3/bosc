package bosc

func init() {
	register("map_tree", func() BinarySearchTree { return newMapTree() })
}

type mapTree struct {
	m map[Comparable]bool
}

func newMapTree() *mapTree {
	t := new(mapTree)
	t.m = make(map[Comparable]bool, 1024)
	return t
}






func (t *mapTree) Add(val Comparable) error { return nil}
func (t *mapTree) Count() uint64 { return 0}
func (t *mapTree) Find(val Comparable) (Comparable, error) {return val, nil}
func (t *mapTree) Min() Comparable {return nil}
func (t *mapTree) Max() Comparable{return nil}
func (t *mapTree) Remove(val Comparable) (found bool) { return false}
func (t *mapTree) Range(valFrom Comparable, valTo Comparable, fn func(node Comparable)) {}
func (t *mapTree) RangeAll(fn func(node Comparable)) {}
func (t *mapTree) RangeFrom(val Comparable, fn func(val Comparable)) {}
func (t *mapTree) RangeTo(val Comparable, fn func(val Comparable)) {}
