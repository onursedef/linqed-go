package linqed_go

// Select projects each element of the collection into a new form
func (q *Queryable[T]) Select(projection func(T) any) *Queryable[T] {
	q.projection = projection
	return q
}
