package lo

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[K comparable, V any](in map[K]V) []K {
	result := make([]K, 0, len(in))
	for k := range in {
		result = append(result, k)
	}
	return result
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[K comparable, V any](in map[K]V) []V {
	result := make([]V, 0, len(in))
	for _, v := range in {
		result = append(result, v)
	}
	return result
}

// Entries transforms a map into a slice of key/value pairs.
func Entries[K comparable, V any](in map[K]V) []Entry[K, V] {
	result := make([]Entry[K, V], 0, len(in))
	for k, v := range in {
		result = append(result, Entry[K, V]{Key: k, Value: v})
	}
	return result
}

// Entry represents a key/value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// FromEntries transforms a slice of key/value pairs into a map.
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	result := make(map[K]V, len(entries))
	for _, e := range entries {
		result[e.Key] = e.Value
	}
	return result
}

// MapValues manipulates a map values and transforms it to a map of another type.
func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(value V, key K) R) map[K]R {
	result := make(map[K]R, len(in))
	for k, v := range in {
		result[k] = iteratee(v, k)
	}
	return result
}

// OmitByKeys returns same map, without blacklisted keys.
func OmitByKeys[K comparable, V any](in map[K]V, keys []K) map[K]V {
	r := make(map[K]V, len(in))
	for k, v := range in {
		r[k] = v
	}
	for _, k := range keys {
		delete(r, k)
	}
	return r
}
