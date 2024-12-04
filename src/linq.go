package linqed_go

import (
	"errors"
	"sort"
)

type Queryable[T any] struct {
	items      []T
	filters    []func(T) bool
	sorter     func(i, j T) bool
	projection func(T) any
	skip       int
	take       int
	hasSkip    bool
	hasTake    bool
}

func From[T any](items []T) *Queryable[T] {
	return &Queryable[T]{items: items}
}

func (q *Queryable[T]) Iterate() []any {
	// First, apply the filters
	var filtered []T
	for _, item := range q.items {
		pass := true
		for _, filter := range q.filters {
			if !filter(item) {
				pass = false
				break
			}
		}

		if pass {
			filtered = append(filtered, item)
		}
	}

	// Then, apply sorting if specified
	if q.sorter != nil {
		sort.SliceStable(filtered, func(i, j int) bool {
			return q.sorter(filtered[i], filtered[j])
		})
	}

	// Apply skip and take logic
	if q.hasSkip && q.skip < len(filtered) {
		filtered = filtered[q.skip:]
	}

	if q.hasTake && q.take < len(filtered) {
		filtered = filtered[:q.take]
	}

	// Then apply the projection
	var result []any
	for _, item := range filtered {
		// If there is a projection, apply it; otherwise, use the item as is.
		if q.projection != nil {
			result = append(result, q.projection(item))
		} else {
			result = append(result, item)
		}
	}

	return result
}

func (q *Queryable[T]) First() (T, error) {
	for _, item := range q.items {
		pass := true
		for _, filter := range q.filters {
			if !filter(item) {
				pass = false
				break
			}
		}

		if pass {
			return item, nil
		}
	}

	var zeroValue T
	return zeroValue, errors.New("no element satisfies the condition")
}

func (q *Queryable[T]) FirstOrDefault() T {
	for _, item := range q.items {
		pass := true
		for _, filter := range q.filters {
			if !filter(item) {
				pass = false
				break
			}
		}

		if pass {
			return item
		}
	}

	var zeroValue T
	return zeroValue
}
