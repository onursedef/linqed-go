package linqed_go

type Joined[T any, U any] struct {
	First  T
	Second U
}

func Join[T, U any, K comparable, R any](q *Queryable[T], other []U, keySelector func(T) K, otherKeySelector func(U) K, resultSelector func(T, U) R) *Queryable[R] {
	var result []R

	for _, item1 := range q.items {
		for _, item2 := range other {
			if keySelector(item1) == otherKeySelector(item2) {
				result = append(result, resultSelector(item1, item2))
			}
		}
	}

	return From(result)
}
