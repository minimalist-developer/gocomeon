package backoff

import (
	"testing"
	"time"
)

func TestBackoff_Next(t *testing.T) {
	backoff := &Backoff{
		InitialBackoff: time.Second,
		MaxBackoff:     10 * time.Second,
		BackoffFactor:  2,
	}
	current := backoff.Next()
	if current != time.Second {
		t.Fatalf("current backoff expected [%s], but [%s] got", time.Second, current)
	}

	current = backoff.Next()
	if current != 2*time.Second {
		t.Fatalf("current backoff expected [%s], but [%s] got", 2*time.Second, current)
	}

	current = backoff.Next()
	if current != 4*time.Second {
		t.Fatalf("current backoff expected [%s], but [%s] got", 4*time.Second, current)
	}

	current = backoff.Next()
	if current != 8*time.Second {
		t.Fatalf("current backoff expected [%s], but [%s] got", 8*time.Second, current)
	}

	current = backoff.Next()
	if current != 10*time.Second {
		t.Fatalf("current backoff expected [%s], but [%s] got", 10*time.Second, current)
	}

	current = backoff.Next()
	if current != 10*time.Second {
		t.Fatalf("current backoff expected [%s], but [%s] got", 10*time.Second, current)
	}

}
