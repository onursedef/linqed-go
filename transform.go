package linqed_go

// Select projects each element of the collection into a new form
func Select[T, R any](q *Queryable[T], projection func(T) R) *Queryable[R] {
	var result []R
	for _, item := range q.items {
		result = append(result, projection(item))
	}
	return From(result)
}
