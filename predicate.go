package lo

// Every returns true if all elements in the collection satisfy the predicate.
func Every[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Some returns true if at least one element in the collection satisfies the predicate.
func Some[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}
	return false
}

// None returns true if no elements in the collection satisfy the predicate.
func None[T any](collection []T, predicate func(T) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return false
		}
	}
	return true
}

// Count returns the number of elements in the collection that satisfy the predicate.
func Count[T any](collection []T, predicate func(T) bool) int {
	count := 0
	for _, item := range collection {
		if predicate(item) {
			count++
		}
	}
	return count
}

// Partition splits the collection into two slices: one with elements that satisfy
// the predicate and one with elements that do not.
func Partition[T any](collection []T, predicate func(T) bool) ([]T, []T) {
	pass := make([]T, 0)
	fail := make([]T, 0)
	for _, item := range collection {
		if predicate(item) {
			pass = append(pass, item)
		} else {
			fail = append(fail, item)
		}
	}
	return pass, fail
}

// CountBy returns a map of keys to counts based on the result of the iteratee function.
func CountBy[T any, K comparable](collection []T, iteratee func(T) K) map[K]int {
	result := make(map[K]int)
	for _, item := range collection {
		key := iteratee(item)
		result[key]++
	}
	return result
}
