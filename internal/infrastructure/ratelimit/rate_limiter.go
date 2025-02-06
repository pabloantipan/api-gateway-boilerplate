package ratelimit

import (
	"sync"
	"time"
)

type RateLimiter struct {
	tokens     map[string]float64   // tokens per client
	lastUpdate map[string]time.Time // last access time per client
	rate       float64              // tokens per second
	capacity   float64              // max tokens
	mu         sync.Mutex
}

func NewRateLimiter(rate, capacity float64) *RateLimiter {
	return &RateLimiter{
		tokens:     make(map[string]float64),
		lastUpdate: make(map[string]time.Time),
		rate:       rate,
		capacity:   capacity,
	}
}

func (rl *RateLimiter) Allow(clientID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	tokens := rl.tokens[clientID]
	lastUpdate := rl.lastUpdate[clientID]

	if !lastUpdate.IsZero() {
		// Add tokens bases on time passed
		elapsed := now.Sub(lastUpdate).Seconds()
		tokens += elapsed * rl.rate
		if tokens > rl.capacity {
			tokens = rl.capacity
		}
	} else {
		// First time access, get full capacity
		tokens = rl.capacity
	}

	// fmt.Printf("Client: %s, Tokens: %f\n", clientID, tokens)

	// Check if client has enough tokens
	if tokens >= 1 {
		tokens--
		rl.tokens[clientID] = tokens
		rl.lastUpdate[clientID] = now
		return true
	}

	return false
}
