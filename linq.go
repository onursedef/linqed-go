package linqed_go

import (
	"errors"
	"sort"
)

// Queryable represents a collection that can be queried
type Queryable[T any] struct {
	// items is the collection
	items []T
	// filters are the filters to apply to the collection
	filters []func(T) bool
	// sorter is the comparator to sort
	sorter func(i, j T) bool
	// projection is the projection function
	projection func(T) any
	// skip is the number of elements to skip
	skip int
	// take is the number of elements to take
	take int
	// hasSkip is true if the skip function has been called
	hasSkip bool
	// hasTake is true if the take function has been called
	hasTake bool
}

// From creates a Queryable from a collection
func From[T any](items []T) *Queryable[T] {
	return &Queryable[T]{items: items}
}

// Iterate returns the items in the collection
func (q *Queryable[T]) Iterate() []any {
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

	if q.sorter != nil {
		sort.SliceStable(filtered, func(i, j int) bool {
			return q.sorter(filtered[i], filtered[j])
		})
	}

	if q.hasSkip && q.skip < len(filtered) {
		filtered = filtered[q.skip:]
	}

	if q.hasTake && q.take < len(filtered) {
		filtered = filtered[:q.take]
	}

	var result []any
	for _, item := range filtered {
		if q.projection != nil {
			result = append(result, q.projection(item))
		} else {
			result = append(result, item)
		}
	}

	return result
}

// First returns the first element in the collection or the error if no element is found
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

// FirstOrDefault returns the first element in the collection or the zero value if no element is found
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
