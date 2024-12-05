package linqed_go

// Concat concatenates the collection with another collection
func (q *Queryable[T]) Concat(other []T) *Queryable[T] {
	q.items = append(q.items, other...)
	return q
}
