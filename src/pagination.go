package linqed_go

func (q *Queryable[T]) Skip(n int) *Queryable[T] {
	q.skip = n
	q.hasSkip = true
	return q
}

func (q *Queryable[T]) Take(n int) *Queryable[T] {
	q.take = n
	q.hasTake = true
	return q
}