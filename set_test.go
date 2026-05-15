package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	result := Uniq([]int{1, 2, 2, 3, 1, 4})
	assert.Equal(t, []int{1, 2, 3, 4}, result)

	empty := Uniq([]string{})
	assert.Equal(t, []string{}, empty)
}

func TestUnion(t *testing.T) {
	result := Union([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)

	single := Union([]string{"a", "b"})
	assert.Equal(t, []string{"a", "b"}, single)
}

func TestIntersection(t *testing.T) {
	result := Intersection([]int{1, 2, 3, 4}, []int{2, 4, 6})
	assert.Equal(t, []int{2, 4}, result)

	none := Intersection([]int{1, 2}, []int{3, 4})
	assert.Equal(t, []int{}, none)
}

func TestDifference(t *testing.T) {
	onlyA, onlyB := Difference([]int{1, 2, 3}, []int{2, 3, 4})
	assert.Equal(t, []int{1}, onlyA)
	assert.Equal(t, []int{4}, onlyB)

	a2, b2 := Difference([]string{"x", "y"}, []string{"y", "z"})
	assert.Equal(t, []string{"x"}, a2)
	assert.Equal(t, []string{"z"}, b2)
}

func TestWithout(t *testing.T) {
	result := Without([]int{1, 2, 3, 4, 5}, 2, 4)
	assert.Equal(t, []int{1, 3, 5}, result)

	none := Without([]string{"a", "b", "c"})
	assert.Equal(t, []string{"a", "b", "c"}, none)
}
