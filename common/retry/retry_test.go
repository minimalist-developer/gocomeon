package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type TestRetryable[T string] struct {
	attempts int
}

func (p *TestRetryable[T]) Required(attempt int, e error) bool {
	return attempt < 5
}

func (p *TestRetryable[T]) DoActionBeforeRetry(attempt int, e error) {
	fmt.Println(fmt.Sprintf("attempt=%d, error=%v", attempt, e))
}

func (p *TestRetryable[T]) DoAction() (T, error) {
	p.attempts++
	return "", errors.New(fmt.Sprintf("do action error %d", p.attempts))
}

func (p *TestRetryable[T]) RetryInterval(attempt int) time.Duration {
	return time.Second
}

func (p *TestRetryable[T]) GetLogOutput() LogOutput {
	return nil
}

func TestInvoke(t *testing.T) {
	retryable := &TestRetryable[string]{}
	result, e := Invoke[string](retryable)
	fmt.Println(fmt.Sprintf("result=%v, error=%v", result, e))
	msg := "do action error 5"
	if e.Error() != msg {
		t.Fatalf("error message expected [%s], but [%s] got", msg, e.Error())
	}
}
