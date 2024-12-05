package linqed_go

// Sum returns the sum of all numeric elements in the collection
func (q *Queryable[T]) Sum() float64 {
	var sum float64
	for _, item := range q.items {
		sum += numeric(item)
	}
	return sum
}

// Count returns the number of elements in the collection
func (q *Queryable[T]) Count() int {
	return len(q.items)
}

// Max returns the maximum value in the collection
func (q *Queryable[T]) Max() (max float64) {
	if len(q.items) > 0 {
		max = numeric(q.items[0])
		for _, item := range q.items {
			if numeric(item) > max {
				max = numeric(item)
			}
		}
	}
	return max
}

// Min returns the minimum value in the collection
func (q *Queryable[T]) Min() (min float64) {
	if len(q.items) > 0 {
		min = numeric(q.items[0])
		for _, item := range q.items {
			if numeric(item) < min {
				min = numeric(item)
			}
		}
	}
	return min
}

// Average returns the average of all numeric elements in the collection
func (q *Queryable[T]) Average() (average float64) {
	sum := q.Sum()
	count := q.Count()

	return sum / float64(count)
}

// Aggregate applies an accumulator function over the collection.
func (q *Queryable[T]) Aggregate(accumulator func(T, T) T) (result T) {
	if len(q.items) == 0 {
		var zero T
		return zero
	}

	result = q.items[0]
	for _, item := range q.items[1:] {
		result = accumulator(result, item)
	}

	return result
}

// All determines whether all elements of a collection satisfy a condition.
func numeric(item any) float64 {
	switch v := any(item).(type) {
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		panic("item is not a numeric type")
	}
}
