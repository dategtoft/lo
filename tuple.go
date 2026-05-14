package lo

// Tuple2 is a group of 2 elements.
type Tuple2[A, B any] struct {
	A A
	B B
}

// Tuple3 is a group of 3 elements.
type Tuple3[A, B, C any] struct {
	A A
	B B
	C C
}

// Tuple4 is a group of 4 elements.
type Tuple4[A, B, C, D any] struct {
	A A
	B B
	C C
	D D
}

// Tuple5 is a group of 5 elements.
type Tuple5[A, B, C, D, E any] struct {
	A A
	B B
	C C
	D D
	E E
}

// T2 creates a tuple from a list of values.
func T2[A, B any](a A, b B) Tuple2[A, B] {
	return Tuple2[A, B]{A: a, B: b}
}

// T3 creates a tuple from a list of values.
func T3[A, B, C any](a A, b B, c C) Tuple3[A, B, C] {
	return Tuple3[A, B, C]{A: a, B: b, C: c}
}

// T4 creates a tuple from a list of values.
func T4[A, B, C, D any](a A, b B, c C, d D) Tuple4[A, B, C, D] {
	return Tuple4[A, B, C, D]{A: a, B: b, C: c, D: d}
}

// T5 creates a tuple from a list of values.
func T5[A, B, C, D, E any](a A, b B, c C, d D, e E) Tuple5[A, B, C, D, E] {
	return Tuple5[A, B, C, D, E]{A: a, B: b, C: c, D: d, E: e}
}

// Unpack2 returns values contained in tuple.
func Unpack2[A, B any](tuple Tuple2[A, B]) (A, B) {
	return tuple.A, tuple.B
}

// Unpack3 returns values contained in tuple.
func Unpack3[A, B, C any](tuple Tuple3[A, B, C]) (A, B, C) {
	return tuple.A, tuple.B, tuple.C
}

// Unpack4 returns values contained in tuple.
func Unpack4[A, B, C, D any](tuple Tuple4[A, B, C, D]) (A, B, C, D) {
	return tuple.A, tuple.B, tuple.C, tuple.D
}

// Unpack5 returns values contained in tuple.
func Unpack5[A, B, C, D, E any](tuple Tuple5[A, B, C, D, E]) (A, B, C, D, E) {
	return tuple.A, tuple.B, tuple.C, tuple.D, tuple.E
}

// Zip2 creates a slice of grouped elements, the first of which contains
// the first elements of the given arrays, the second contains the second
// elements, and so on.
// If one array is shorter than the other, its missing elements are zero values.
func Zip2[A, B any](a []A, b []B) []Tuple2[A, B] {
	size := max(len(a), len(b))
	result := make([]Tuple2[A, B], size)
	for i := 0; i < size; i++ {
		var va A
		var vb B
		if i < len(a) {
			va = a[i]
		}
		if i < len(b) {
			vb = b[i]
		}
		result[i] = T2(va, vb)
	}
	return result
}

// Unzip2 accepts a slice of grouped elements and creates a slice of the first
// elements and a slice of the second elements.
func Unzip2[A, B any](collection []Tuple2[A, B]) ([]A, []B) {
	a := make([]A, len(collection))
	b := make([]B, len(collection))
	for i, t := range collection {
		a[i] = t.A
		b[i] = t.B
	}
	return a, b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
