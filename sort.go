package linqed_go

// OrderBy sorts the collection based on a comparator
func (q *Queryable[T]) OrderBy(comparator func(i, j T) bool) *Queryable[T] {
	q.sorter = comparator
	return q
}
