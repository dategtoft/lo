package lo

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetry_Success(t *testing.T) {
	is := assert.New(t)

	calls := 0
	failed, err := Retry(RetryConfig{Attempts: 3, Delay: 0, BackoffFactor: 1}, func(attempt int) error {
		calls++
		return nil
	})

	is.NoError(err)
	is.Equal(0, failed)
	is.Equal(1, calls)
}

func TestRetry_EventualSuccess(t *testing.T) {
	is := assert.New(t)

	calls := 0
	failed, err := Retry(RetryConfig{Attempts: 5, Delay: 0, BackoffFactor: 1}, func(attempt int) error {
		calls++
		if attempt < 2 {
			return errors.New("not yet")
		}
		return nil
	})

	is.NoError(err)
	is.Equal(2, failed)
	is.Equal(3, calls)
}

func TestRetry_AllFail(t *testing.T) {
	is := assert.New(t)

	sentinel := errors.New("always fail")
	calls := 0
	failed, err := Retry(RetryConfig{Attempts: 3, Delay: 0, BackoffFactor: 1}, func(attempt int) error {
		calls++
		return sentinel
	})

	is.ErrorIs(err, sentinel)
	is.Equal(3, failed)
	is.Equal(3, calls)
}

func TestRetryWithResult_Success(t *testing.T) {
	is := assert.New(t)

	result, failed, err := RetryWithResult(RetryConfig{Attempts: 3, Delay: 0, BackoffFactor: 1}, func(attempt int) (string, error) {
		if attempt < 1 {
			return "", errors.New("not ready")
		}
		return "ok", nil
	})

	is.NoError(err)
	is.Equal("ok", result)
	is.Equal(1, failed)
}

func TestRetry_BackoffFactor(t *testing.T) {
	is := assert.New(t)

	start := time.Now()
	cfg := RetryConfig{
		Attempts:      3,
		Delay:         10 * time.Millisecond,
		BackoffFactor: 2,
		MaxDelay:      50 * time.Millisecond,
	}
	_, _ = Retry(cfg, func(attempt int) error {
		return errors.New("fail")
	})
	elapsed := time.Since(start)
	// 10ms + 20ms = 30ms minimum between attempts
	is.GreaterOrEqual(elapsed, 25*time.Millisecond)
}
