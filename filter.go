package linqed_go

// Where filters the collection based on a predicate
func (q *Queryable[T]) Where(predicate func(T) bool) *Queryable[T] {
	q.filters = append(q.filters, predicate)
	return q
}

// Any returns true if any element in the collection satisfies the predicate
func (q *Queryable[T]) Any() bool {
	for _, item := range q.items {
		pass := true
		for _, filter := range q.filters {
			if !filter(item) {
				pass = false
				break
			}
		}
		if pass {
			return true
		}
	}

	return false
}

// All returns true if all elements in the collection satisfy the predicate
func (q *Queryable[T]) All(predicate func(T) bool) bool {
	for _, item := range q.items {
		if !predicate(item) {
			return false
		}
	}

	return true
}
