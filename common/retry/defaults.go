package retry

import (
	"errors"
	"time"
)

type ErrorActionIsNil struct {
}

func (p *ErrorActionIsNil) Error() string {
	return "Action can not be nil"
}

// Action Do action.
type Action[T any] func() (T, error)

// ActionBeforeRetry Do something before retry again.
type ActionBeforeRetry func(int, error)

// RetryableInterval Retry interval function
type RetryableInterval func(int) time.Duration

type RetryableTimes[T any] struct {
	// Attempts the number retry attempts
	Attempts int
	// RetryableInterval retry time interval, unit is millisecond
	RetryableInterval RetryableInterval
	// Action do action
	Action Action[T]
	// ActionBeforeRetry do something before retry again
	ActionBeforeRetry ActionBeforeRetry
	// LogOutput print log
	LogOutput LogOutput
}

func (p *RetryableTimes[T]) Required(attempt int, e error) bool {
	if errors.Is(e, &ErrorActionIsNil{}) {
		return false
	}
	return attempt < p.Attempts
}

func (p *RetryableTimes[T]) RetryInterval(attempt int) time.Duration {
	return p.RetryableInterval(attempt)
}

func (p *RetryableTimes[T]) DoActionBeforeRetry(attempt int, e error) {
	if p.ActionBeforeRetry != nil {
		p.ActionBeforeRetry(attempt, e)
	}
}

func (p *RetryableTimes[T]) DoAction() (T, error) {
	if p.Action == nil {
		var result T
		return result, &ErrorActionIsNil{}
	}
	return p.Action()
}

func (p *RetryableTimes[T]) GetLogOutput() LogOutput {
	return p.LogOutput
}
