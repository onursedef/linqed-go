package linqed_go

func (q *Queryable[T]) Where(predicate func(T) bool) *Queryable[T] {
	q.filters = append(q.filters, predicate)
	return q
}

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

func (q *Queryable[T]) All(predicate func(T) bool) bool {
	for _, item := range q.items {
		if !predicate(item) {
			return false
		}
	}

	return true
}
