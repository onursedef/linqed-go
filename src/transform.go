package linqed_go

func (q *Queryable[T]) Select(projection func(T) any) *Queryable[T] {
	q.projection = projection
	return q
}
