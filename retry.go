package lo

import (
	"time"
)

// RetryConfig holds configuration for retry behaviour.
type RetryConfig struct {
	// Attempts is the maximum number of attempts (including the first call).
	Attempts int
	// Delay is the duration to wait between attempts.
	Delay time.Duration
	// MaxDelay caps the exponential backoff delay. 0 means no cap.
	MaxDelay time.Duration
	// BackoffFactor multiplies the delay after each failed attempt.
	// A value of 1 means constant delay.
	BackoffFactor float64
}

// DefaultRetryConfig is a sensible default retry configuration.
var DefaultRetryConfig = RetryConfig{
	Attempts:      3,
	Delay:         100 * time.Millisecond,
	MaxDelay:      0,
	BackoffFactor: 1,
}

// Retry calls the function f up to config.Attempts times, stopping on success.
// It returns the number of failed attempts and any final error.
func Retry(config RetryConfig, f func(attempt int) error) (int, error) {
	var err error
	delay := config.Delay
	if config.BackoffFactor == 0 {
		config.BackoffFactor = 1
	}
	for attempt := 0; attempt < config.Attempts; attempt++ {
		if err = f(attempt); err == nil {
			return attempt, nil
		}
		if attempt < config.Attempts-1 && delay > 0 {
			time.Sleep(delay)
			newDelay := time.Duration(float64(delay) * config.BackoffFactor)
			if config.MaxDelay > 0 && newDelay > config.MaxDelay {
				newDelay = config.MaxDelay
			}
			delay = newDelay
		}
	}
	return config.Attempts, err
}

// RetryWithResult calls f up to config.Attempts times and returns the result
// of the first successful call, or the zero value and final error.
func RetryWithResult[T any](config RetryConfig, f func(attempt int) (T, error)) (T, int, error) {
	var (
		result T
		err    error
	)
	delay := config.Delay
	if config.BackoffFactor == 0 {
		config.BackoffFactor = 1
	}
	for attempt := 0; attempt < config.Attempts; attempt++ {
		if result, err = f(attempt); err == nil {
			return result, attempt, nil
		}
		if attempt < config.Attempts-1 && delay > 0 {
			time.Sleep(delay)
			newDelay := time.Duration(float64(delay) * config.BackoffFactor)
			if config.MaxDelay > 0 && newDelay > config.MaxDelay {
				newDelay = config.MaxDelay
			}
			delay = newDelay
		}
	}
	return result, config.Attempts, err
}
