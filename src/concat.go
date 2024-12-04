package linqed_go

func (q *Queryable[T]) Concat(other []T) *Queryable[T] {
	q.items = append(q.items, other...)
	return q
}
