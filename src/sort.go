package linqed_go

func (q *Queryable[T]) OrderBy(comparator func(i, j T) bool) *Queryable[T] {
	q.sorter = comparator
	return q
}
