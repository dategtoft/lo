package lo

import (
	"sync"
	"time"
)

// Generator creates a channel that emits values produced by the generator function.
// The channel is closed when the generator returns false.
func Generator[T any](bufferSize int, generator func(yield func(T))) <-chan T {
	ch := make(chan T, bufferSize)
	go func() {
		defer close(ch)
		generator(func(v T) {
			ch <- v
		})
	}()
	return ch
}

// Batch collects items from a channel into slices of the given size.
// The last batch may be smaller than size.
func Batch[T any](ch <-chan T, size int) <-chan []T {
	out := make(chan []T)
	go func() {
		defer close(out)
		batch := make([]T, 0, size)
		for v := range ch {
			batch = append(batch, v)
			if len(batch) == size {
				out <- batch
				batch = make([]T, 0, size)
			}
		}
		if len(batch) > 0 {
			out <- batch
		}
	}()
	return out
}

// BatchWithTimeout collects items from a channel into slices, flushing when
// the batch reaches size or the timeout elapses.
func BatchWithTimeout[T any](ch <-chan T, size int, timeout time.Duration) <-chan []T {
	out := make(chan []T)
	go func() {
		defer close(out)
		batch := make([]T, 0, size)
		ticker := time.NewTicker(timeout)
		defer ticker.Stop()
		flush := func() {
			if len(batch) > 0 {
				out <- batch
				batch = make([]T, 0, size)
			}
		}
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					flush()
					return
				}
				batch = append(batch, v)
				if len(batch) >= size {
					flush()
					ticker.Reset(timeout)
				}
			case <-ticker.C:
				flush()
			}
		}
	}()
	return out
}

// Merge merges multiple channels into a single channel.
func Merge[T any](channels ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan T) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// ChannelToSlice drains a channel into a slice.
func ChannelToSlice[T any](ch <-chan T) []T {
	result := make([]T, 0)
	for v := range ch {
		result = append(result, v)
	}
	return result
}
