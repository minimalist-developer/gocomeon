package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jinjianfeng-chn/gocomeon/common/backoff"
)

func TestRetryableTimes(t *testing.T) {

	retryableTimes := RetryableTimes[string]{
		Attempts: 3,
		RetryableInterval: func(i int) time.Duration {
			return time.Second
		},
		Action: func() (string, error) {
			return "", errors.New("do action error")
		},
	}

	t.Run("normal", func(t *testing.T) {
		r, e := Invoke[string](&retryableTimes)
		fmt.Println(r)

		msg := "do action error"
		if e.Error() != msg {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", msg, e.Error()))
		}
	})

	t.Run("action is nil", func(t *testing.T) {
		action := retryableTimes.Action
		retryableTimes.Action = nil
		_, e := Invoke[string](&retryableTimes)

		if !errors.Is(e, &ErrorActionIsNil{}) {
			t.Fatal(fmt.Sprintf("error message expected [%s], but [%s] got", &ErrorActionIsNil{}, e.Error()))
		}
		retryableTimes.Action = action
	})

	t.Run("test backoff retry", func(t *testing.T) {
		retryableInterval := retryableTimes.RetryableInterval
		bBackoff := &backoff.Backoff{
			InitialBackoff: time.Second,
			MaxBackoff:     10 * time.Second,
			BackoffFactor:  2,
		}
		retryableTimes.RetryableInterval = func(i int) time.Duration {
			return bBackoff.Next()
		}
		beginTime := time.Now().UnixMilli()
		_, _ = Invoke[string](&retryableTimes)
		now := time.Now().UnixMilli()
		if now-beginTime < 3000 || now-beginTime > 4000 {
			t.Fatal(fmt.Sprintf("The expected running time is between 3 and 4 seconds, but [%d] ms got", now-beginTime))
		}
		retryableTimes.RetryableInterval = retryableInterval
	})
}
