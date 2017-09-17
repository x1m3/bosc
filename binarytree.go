package bosc


type Comparable interface {
	Compare(b Comparable) int
	Key() interface{}
}

type BinarySearchTree interface {
	Add(val Comparable) error
	Count() uint64
	Find(val Comparable) (Comparable, error)
	Min() Comparable
	Max() Comparable
	Remove(val Comparable) (found bool)
	Range(valFrom Comparable, valTo Comparable, fn func(node Comparable))
	RangeAll(fn func(node Comparable))
	RangeFrom(val Comparable, fn func(val Comparable))
	RangeTo(val Comparable, fn func(val Comparable))
}


