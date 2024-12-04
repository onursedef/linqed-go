package linqed_go

import "fmt"

func (q *Queryable[T]) Distinct() *Queryable[T] {
	seen := make(map[string]bool)
	var result []T
	for _, item := range q.items {
		key := fmt.Sprintf("%v", item)
		if !seen[key] {
			seen[key] = true
			result = append(result, item)
		}
	}

	q.items = result
	return q
}
