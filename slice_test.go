package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	result := Filter([]int{1, 2, 3, 4, 5}, func(item int, _ int) bool {
		return item%2 == 0
	})
	assert.Equal(t, []int{2, 4}, result)
}

func TestMap(t *testing.T) {
	result := Map([]int{1, 2, 3}, func(item int, _ int) string {
		return fmt.Sprintf("%d", item*2)
	})
	assert.Equal(t, []string{"2", "4", "6"}, result)
}

func TestReduce(t *testing.T) {
	result := Reduce([]int{1, 2, 3, 4, 5}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	assert.Equal(t, 15, result)
}

func TestForEach(t *testing.T) {
	sum := 0
	ForEach([]int{1, 2, 3}, func(item int, _ int) {
		sum += item
	})
	assert.Equal(t, 6, sum)
}

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{1, 2, 3}, 2))
	assert.False(t, Contains([]int{1, 2, 3}, 5))
}

func TestUniq(t *testing.T) {
	result := Uniq([]int{1, 2, 2, 3, 3, 3})
	assert.Equal(t, []int{1, 2, 3}, result)
}

func TestFlatten(t *testing.T) {
	result := Flatten([][]int{{1, 2}, {3, 4}, {5}})
	assert.Equal(t, []int{1, 2, 3, 4, 5}, result)
}
