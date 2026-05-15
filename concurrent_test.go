package lo

import (
	"sort"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := ParallelMap(input, func(item int, index int) int {
		return item * 2
	})
	assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
}

func TestParallelMap_PreservesOrder(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	result := ParallelMap(input, func(item string, index int) string {
		return item + item
	})
	assert.Equal(t, []string{"aa", "bb", "cc", "dd"}, result)
}

func TestParallelForEach(t *testing.T) {
	var counter int64
	input := []int{1, 2, 3, 4, 5}
	ParallelForEach(input, func(item int, index int) {
		atomic.AddInt64(&counter, int64(item))
	})
	assert.Equal(t, int64(15), counter)
}

func TestParallelFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	result := ParallelFilter(input, func(item int, index int) bool {
		return item%2 == 0
	})
	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestParallelFilter_PreservesOrder(t *testing.T) {
	input := []int{5, 3, 1, 4, 2}
	result := ParallelFilter(input, func(item int, index int) bool {
		return item > 2
	})
	assert.Equal(t, []int{5, 3, 4}, result)
}

func TestParallelGroupBy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	result := ParallelGroupBy(input, func(item int) string {
		if item%2 == 0 {
			return "even"
		}
		return "odd"
	})

	assert.Len(t, result["even"], 3)
	assert.Len(t, result["odd"], 3)

	sort.Ints(result["even"])
	sort.Ints(result["odd"])
	assert.Equal(t, []int{2, 4, 6}, result["even"])
	assert.Equal(t, []int{1, 3, 5}, result["odd"])
}

func TestParallelMap_Empty(t *testing.T) {
	result := ParallelMap([]int{}, func(item int, index int) int {
		return item
	})
	assert.Empty(t, result)
}
