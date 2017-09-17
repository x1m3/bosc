package bosc

//https://es.wikipedia.org/wiki/%C3%81rbol_binario_de_b%C3%BAsqueda#Tipos_de_.C3.A1rboles_binarios_de_b.C3.BAsqueda

type Comparable interface {
	Compare(b Comparable) int
	Key() interface{} // Probablemente no es necesario
}

type BinarySearchTree interface {
	Add(val Comparable) error
	Count() uint64
	Find(val Comparable) (Comparable, error)
	Min() *Comparable
	Max() *Comparable
	Remove(val Comparable) (found bool)
	Range(valFrom Comparable, valTo Comparable, fn func(node Comparable))
	RangeAll(fn func(node Comparable))
	RangeFrom(val Comparable, fn func(val Comparable))
	RangeTo(val Comparable, fn func(val Comparable))


}


