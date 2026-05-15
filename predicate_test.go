package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvery(t *testing.T) {
	is := assert.New(t)

	is.True(Every([]int{2, 4, 6, 8}, func(n int) bool { return n%2 == 0 }))
	is.False(Every([]int{2, 3, 6, 8}, func(n int) bool { return n%2 == 0 }))
	is.True(Every([]int{}, func(n int) bool { return n%2 == 0 }))
}

func TestSomeP(t *testing.T) {
	is := assert.New(t)

	is.True(Some([]int{1, 2, 3}, func(n int) bool { return n%2 == 0 }))
	is.False(Some([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 }))
	is.False(Some([]int{}, func(n int) bool { return n%2 == 0 }))
}

func TestNone(t *testing.T) {
	is := assert.New(t)

	is.True(None([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 }))
	is.False(None([]int{1, 2, 3}, func(n int) bool { return n%2 == 0 }))
	is.True(None([]int{}, func(n int) bool { return n%2 == 0 }))
}

func TestCount(t *testing.T) {
	is := assert.New(t)

	is.Equal(3, Count([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 }))
	is.Equal(0, Count([]int{1, 3, 5}, func(n int) bool { return n%2 == 0 }))
	is.Equal(0, Count([]int{}, func(n int) bool { return n%2 == 0 }))
}

func TestPartition(t *testing.T) {
	is := assert.New(t)

	pass, fail := Partition([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	is.Equal([]int{2, 4, 6}, pass)
	is.Equal([]int{1, 3, 5}, fail)

	pass2, fail2 := Partition([]int{}, func(n int) bool { return n > 0 })
	is.Empty(pass2)
	is.Empty(fail2)
}

func TestCountBy(t *testing.T) {
	is := assert.New(t)

	result := CountBy([]string{"foo", "bar", "baz", "foo", "foo", "bar"}, func(s string) string { return s })
	is.Equal(map[string]int{"foo": 3, "bar": 2, "baz": 1}, result)

	empty := CountBy([]int{}, func(n int) int { return n })
	is.Empty(empty)
}
