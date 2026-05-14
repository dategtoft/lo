package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 15, Sum([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, Sum([]int{}))
	assert.InDelta(t, 6.5, Sum([]float64{1.5, 2.0, 3.0}), 0.0001)
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min([]int{3, 1, 4, 1, 5, 9}))
	assert.Equal(t, 0, Min([]int{}))
	assert.Equal(t, "apple", Min([]string{"banana", "apple", "cherry"}))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 9, Max([]int{3, 1, 4, 1, 5, 9}))
	assert.Equal(t, 0, Max([]int{}))
	assert.Equal(t, "cherry", Max([]string{"banana", "apple", "cherry"}))
}

func TestClamp(t *testing.T) {
	assert.Equal(t, 5, Clamp(5, 1, 10))
	assert.Equal(t, 1, Clamp(-3, 1, 10))
	assert.Equal(t, 10, Clamp(15, 1, 10))
	assert.InDelta(t, 0.5, Clamp(0.5, 0.0, 1.0), 0.0001)
	assert.InDelta(t, 0.0, Clamp(-0.5, 0.0, 1.0), 0.0001)
	assert.InDelta(t, 1.0, Clamp(1.5, 0.0, 1.0), 0.0001)
}

func TestMean(t *testing.T) {
	assert.InDelta(t, 3.0, Mean([]int{1, 2, 3, 4, 5}), 0.0001)
	assert.InDelta(t, 0.0, Mean([]int{}), 0.0001)
	assert.InDelta(t, 2.5, Mean([]float64{1.0, 2.0, 3.0, 4.0}), 0.0001)
}
