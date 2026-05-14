package lo

// Filter iterates over a collection and returns an array of all
// elements for which the predicate function returns true.
func Filter[V any](collection []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(collection))
	for i, item := range collection {
		if predicate(item, i) {
			result = append(result, item)
		}
	}
	return result
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))
	for i, item := range collection {
		result[i] = iteratee(item, i)
	}
	return result
}

// Reduce reduces a collection to a single value using an accumulator function.
func Reduce[T any, R any](collection []T, accumulator func(agg R, item T, index int) R, initial R) R {
	result := initial
	for i, item := range collection {
		result = accumulator(result, item, i)
	}
	return result
}

// ForEach iterates over elements of a collection and invokes the function
// for each element.
func ForEach[T any](collection []T, iteratee func(item T, index int)) {
	for i, item := range collection {
		iteratee(item, i)
	}
}

// Contains returns true if an element is present in a collection.
func Contains[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// Uniq returns a duplicate-free version of an array, in which only the
// first occurrence of each element is kept.
func Uniq[T comparable](collection []T) []T {
	seen := make(map[T]struct{}, len(collection))
	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, ok := seen[item]; !ok {
			seen[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Flatten returns an array a single level deep.
func Flatten[T any](collection [][]T) []T {
	result := make([]T, 0)
	for _, inner := range collection {
		result = append(result, inner...)
	}
	return result
}
