package lo

// Uniq returns a slice with duplicate values removed.
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

// Union returns a new slice containing all unique elements from all provided slices.
func Union[T comparable](collections ...[]T) []T {
	seen := make(map[T]struct{})
	result := []T{}
	for _, collection := range collections {
		for _, item := range collection {
			if _, ok := seen[item]; !ok {
				seen[item] = struct{}{}
				result = append(result, item)
			}
		}
	}
	return result
}

// Intersection returns a slice of elements that are present in all provided slices.
func Intersection[T comparable](a []T, b []T) []T {
	seen := make(map[T]struct{}, len(b))
	for _, item := range b {
		seen[item] = struct{}{}
	}
	result := []T{}
	for _, item := range a {
		if _, ok := seen[item]; ok {
			result = append(result, item)
			delete(seen, item)
		}
	}
	return result
}

// Difference returns two slices: elements only in a, and elements only in b.
func Difference[T comparable](a []T, b []T) ([]T, []T) {
	inB := make(map[T]struct{}, len(b))
	for _, item := range b {
		inB[item] = struct{}{}
	}
	inA := make(map[T]struct{}, len(a))
	for _, item := range a {
		inA[item] = struct{}{}
	}

	onlyA := []T{}
	for _, item := range a {
		if _, ok := inB[item]; !ok {
			onlyA = append(onlyA, item)
		}
	}
	onlyB := []T{}
	for _, item := range b {
		if _, ok := inA[item]; !ok {
			onlyB = append(onlyB, item)
		}
	}
	return onlyA, onlyB
}

// Without returns a slice excluding all given values.
func Without[T comparable](collection []T, exclude ...T) []T {
	excluded := make(map[T]struct{}, len(exclude))
	for _, item := range exclude {
		excluded[item] = struct{}{}
	}
	result := make([]T, 0, len(collection))
	for _, item := range collection {
		if _, ok := excluded[item]; !ok {
			result = append(result, item)
		}
	}
	return result
}
