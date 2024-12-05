package linqed_go

// Skip skips the first n elements in the collection
func (q *Queryable[T]) Skip(n int) *Queryable[T] {
	q.skip = n
	q.hasSkip = true
	return q
}

// Take takes the first n elements in the collection
func (q *Queryable[T]) Take(n int) *Queryable[T] {
	q.take = n
	q.hasTake = true
	return q
}
