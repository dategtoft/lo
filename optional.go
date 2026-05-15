package lo

// Option represents an optional value: every Option is either Some and contains
// a value, or None and does not. Option types are useful for representing
// absence of values without using nil pointers or sentinel values.
type Option[T any] struct {
	value    T
	hasValue bool
}

// Some creates an Option containing the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{
		value:    value,
		hasValue: true,
	}
}

// None creates an empty Option with no value.
func None[T any]() Option[T] {
	return Option[T]{
		hasValue: false,
	}
}

// IsPresent returns true if the Option contains a value.
func (o Option[T]) IsPresent() bool {
	return o.hasValue
}

// IsAbsent returns true if the Option does not contain a value.
func (o Option[T]) IsAbsent() bool {
	return !o.hasValue
}

// Get returns the value and a boolean indicating whether the value is present.
func (o Option[T]) Get() (T, bool) {
	return o.value, o.hasValue
}

// MustGet returns the value if present, or panics if the Option is empty.
func (o Option[T]) MustGet() T {
	if !o.hasValue {
		panic("lo.Option: called MustGet on a None value")
	}
	return o.value
}

// OrElse returns the contained value if present, otherwise returns the
// provided default value.
func (o Option[T]) OrElse(defaultValue T) T {
	if o.hasValue {
		return o.value
	}
	return defaultValue
}

// OrElseGet returns the contained value if present, otherwise calls the
// provided function and returns its result.
func (o Option[T]) OrElseGet(fn func() T) T {
	if o.hasValue {
		return o.value
	}
	return fn()
}

// Map applies the given function to the value inside the Option if present,
// returning a new Option containing the result. If the Option is empty,
// returns None.
func MapOption[T any, R any](o Option[T], fn func(T) R) Option[R] {
	if o.hasValue {
		return Some(fn(o.value))
	}
	return None[R]()
}

// FlatMap applies the given function to the value inside the Option if present,
// returning the Option produced by the function. If the Option is empty,
// returns None.
func FlatMapOption[T any, R any](o Option[T], fn func(T) Option[R]) Option[R] {
	if o.hasValue {
		return fn(o.value)
	}
	return None[R]()
}

// Filter returns the Option if it contains a value and the predicate returns
// true, otherwise returns None.
func (o Option[T]) Filter(predicate func(T) bool) Option[T] {
	if o.hasValue && predicate(o.value) {
		return o
	}
	return None[T]()
}

// ToSlice returns a slice containing the value if present, or an empty slice.
func (o Option[T]) ToSlice() []T {
	if o.hasValue {
		return []T{o.value}
	}
	return []T{}
}

// FromPtr creates an Option from a pointer. If the pointer is nil, returns None;
// otherwise returns Some with the dereferenced value.
func FromPtr[T any](ptr *T) Option[T] {
	if ptr == nil {
		return None[T]()
	}
	return Some(*ptr)
}

// ToPtr returns a pointer to the value if present, or nil if absent.
func (o Option[T]) ToPtr() *T {
	if o.hasValue {
		v := o.value
		return &v
	}
	return nil
}
