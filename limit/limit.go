package middleware

import (
	"time"
)

type Limiter interface {
	Allow() bool
}

// 基于 Token 的限流器
type tokenLimiter struct {
	rate   int
	bucket chan struct{}
	block  bool
}

func (t *tokenLimiter) Allow() bool {
	if t.block {
		return t.Acquire()
	}
	return t.TryAcquire()

}

func (t *tokenLimiter) TryAcquire() bool {
	select {
	case <-t.bucket:
		return true
	default:
		return false
	}
}

func (t *tokenLimiter) Acquire() bool {
	<-t.bucket
	return true
}

func NewTokenLimiter(rate int, block bool) *tokenLimiter {
	bucket := make(chan struct{}, rate)
	for {
		if len(bucket) >= rate {
			break
		}
		bucket <- struct{}{}
	}
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(1000/rate))
			if len(bucket) < rate {
				bucket <- struct{}{}
			}
		}
	}()
	return &tokenLimiter{
		rate:   rate,
		bucket: bucket,
		block:  block,
	}
}
