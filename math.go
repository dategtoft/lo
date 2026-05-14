package lo

import "golang.org/x/exp/constraints"

// Sum returns the sum of all elements in the collection.
func Sum[T constraints.Number](collection []T) T {
	var sum T
	for _, v := range collection {
		sum += v
	}
	return sum
}

// Min returns the minimum value in the collection.
// Returns zero value if the collection is empty.
func Min[T constraints.Ordered](collection []T) T {
	var min T
	if len(collection) == 0 {
		return min
	}
	min = collection[0]
	for _, v := range collection[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value in the collection.
// Returns zero value if the collection is empty.
func Max[T constraints.Ordered](collection []T) T {
	var max T
	if len(collection) == 0 {
		return max
	}
	max = collection[0]
	for _, v := range collection[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Clamp restricts a value to be within the range [min, max].
func Clamp[T constraints.Ordered](value, min, max T) T {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Mean returns the arithmetic mean of the collection.
// Returns zero value if the collection is empty.
func Mean[T constraints.Number](collection []T) float64 {
	if len(collection) == 0 {
		return 0
	}
	var sum float64
	for _, v := range collection {
		sum += float64(v)
	}
	return sum / float64(len(collection))
}
