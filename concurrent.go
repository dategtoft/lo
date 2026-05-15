package lo

import (
	"sync"
)

// ParallelMap applies the given function to each element of the slice concurrently
// and returns a new slice with the results.
func ParallelMap[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))
	var wg sync.WaitGroup

	for i, item := range collection {
		wg.Add(1)
		go func(idx int, val T) {
			defer wg.Done()
			result[idx] = iteratee(val, idx)
		}(i, item)
	}

	wg.Wait()
	return result
}

// ParallelForEach applies the given function to each element of the slice concurrently.
func ParallelForEach[T any](collection []T, iteratee func(item T, index int)) {
	var wg sync.WaitGroup

	for i, item := range collection {
		wg.Add(1)
		go func(idx int, val T) {
			defer wg.Done()
			iteratee(val, idx)
		}(i, item)
	}

	wg.Wait()
}

// ParallelFilter applies the given predicate to each element of the slice concurrently
// and returns a new slice containing only elements for which the predicate returned true.
// Order of results is preserved.
func ParallelFilter[T any](collection []T, predicate func(item T, index int) bool) []T {
	keep := make([]bool, len(collection))
	var wg sync.WaitGroup

	for i, item := range collection {
		wg.Add(1)
		go func(idx int, val T) {
			defer wg.Done()
			keep[idx] = predicate(val, idx)
		}(i, item)
	}

	wg.Wait()

	result := make([]T, 0, len(collection))
	for i, item := range collection {
		if keep[i] {
			result = append(result, item)
		}
	}
	return result
}

// ParallelGroupBy applies the given key function to each element of the slice concurrently
// and returns a map grouping elements by their key. Order within each group is not guaranteed.
func ParallelGroupBy[T any, K comparable](collection []T, keyFn func(item T) K) map[K][]T {
	type indexedItem struct {
		index int
		key   K
		val   T
	}

	results := make([]indexedItem, len(collection))
	var wg sync.WaitGroup

	for i, item := range collection {
		wg.Add(1)
		go func(idx int, val T) {
			defer wg.Done()
			results[idx] = indexedItem{index: idx, key: keyFn(val), val: val}
		}(i, item)
	}

	wg.Wait()

	grouped := make(map[K][]T)
	for _, r := range results {
		grouped[r.key] = append(grouped[r.key], r.val)
	}
	return grouped
}
