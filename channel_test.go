package lo

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	is := assert.New(t)

	ch := Generator(0, func(yield func(int)) {
		for i := 0; i < 5; i++ {
			yield(i)
		}
	})

	result := ChannelToSlice(ch)
	is.Equal([]int{0, 1, 2, 3, 4}, result)
}

func TestBatch(t *testing.T) {
	is := assert.New(t)

	ch := Generator(0, func(yield func(int)) {
		for i := 1; i <= 7; i++ {
			yield(i)
		}
	})

	batches := ChannelToSlice(Batch(ch, 3))
	is.Len(batches, 3)
	is.Equal([]int{1, 2, 3}, batches[0])
	is.Equal([]int{4, 5, 6}, batches[1])
	is.Equal([]int{7}, batches[2])
}

func TestBatchWithTimeout(t *testing.T) {
	is := assert.New(t)

	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	batches := ChannelToSlice(BatchWithTimeout(ch, 10, 50*time.Millisecond))
	is.Len(batches, 1)
	is.Equal([]int{1, 2, 3}, batches[0])
}

func TestMerge(t *testing.T) {
	is := assert.New(t)

	ch1 := Generator(0, func(yield func(int)) {
		yield(1)
		yield(2)
	})
	ch2 := Generator(0, func(yield func(int)) {
		yield(3)
		yield(4)
	})

	result := ChannelToSlice(Merge(ch1, ch2))
	sort.Ints(result)
	is.Equal([]int{1, 2, 3, 4}, result)
}

func TestChannelToSlice(t *testing.T) {
	is := assert.New(t)

	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	close(ch)

	result := ChannelToSlice(ch)
	is.Equal([]string{"a", "b", "c"}, result)
}
