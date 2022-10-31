package backoff

import "time"

type Backoff struct {
	// current current time
	current time.Duration
	// InitialBackoff the backoff interval for the first retry attempt
	InitialBackoff time.Duration
	// MaxBackoff the maximum backoff interval for any retry attempt
	MaxBackoff time.Duration
	// BackoffFactor the factor by which the InitialBackoff is exponentially
	// incremented for consecutive retry attempts.
	// For example, a backoff factor of 2.5 will result in a backoff of
	// InitialBackoff * 2.5 * 2.5 on the second attempt.
	BackoffFactor float64
}

func (p *Backoff) Next() time.Duration {
	if p.current < p.InitialBackoff {
		p.current = p.InitialBackoff
	} else {
		c := float64(p.current) * p.BackoffFactor
		if c > float64(p.MaxBackoff) {
			c = float64(p.MaxBackoff)
		}
		p.current = time.Duration(c)
	}
	return p.current
}

func (p *Backoff) Current() time.Duration {
	return p.current
}
