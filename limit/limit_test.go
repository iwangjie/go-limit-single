package middleware

import (
	"fmt"
	"testing"
	"time"
)

var limiter = NewTokenLimiter(5, false)
var blockLimiter = NewTokenLimiter(5, true)

func TestTokenLimiter(t *testing.T) {
	limiter := NewTokenLimiter(5, false)
	go func() {
		for i := 0; i < 10; i++ {
			if limiter.Allow() {
				fmt.Println("Allow", time.Now())
			} else {
				fmt.Println("Deny")
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			if limiter.Allow() {
				fmt.Println("Allow", time.Now())
			} else {
				fmt.Println("Deny")
			}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			if limiter.Allow() {
				fmt.Println("Allow", time.Now())
			} else {
				fmt.Println("Deny")
			}
		}
	}()
	select {}
}

func BenchmarkTokenLimiter_Allow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		limiter.Allow()
	}
}

func BenchmarkBlockTokenLimiter_Allow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blockLimiter.Allow()
	}
}
