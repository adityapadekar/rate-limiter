package tokenbucketinmemory

import (
	"log/slog"
	"sync"
	"time"
)

type service struct {
	data map[string]map[string]*bucket
	sync.Mutex
}

type Service interface {
	RateLimitRequest(ipAddr, pattern string) bool
}

func NewService() Service {
	return &service{
		data: make(map[string]map[string]*bucket),
	}
}

func (s *service) refill(rateLimitBucket *bucket, token int) bool {
	t := time.Since(rateLimitBucket.lastRefillTime)

	if t >= time.Second {
		rateLimitBucket.token = token
		rateLimitBucket.lastRefillTime = time.Now()
		return true
	}

	return false
}

func (s *service) RateLimitRequest(ipAddr, pattern string) bool {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.data[ipAddr]; !ok {
		slog.Info("no entry found", "ipAddr", ipAddr)
		s.data[ipAddr] = make(map[string]*bucket)
	}

	token := rateLimiterConfig[pattern].token
	if _, ok := s.data[ipAddr][pattern]; !ok {
		slog.Info("no bucket found", "pattern", pattern)
		b := bucket{
			token:          token,
			lastRefillTime: time.Now(),
		}
		s.data[ipAddr][pattern] = &b
	}

	rateLimitBucket := s.data[ipAddr][pattern]
	refillSuccess := true
	if rateLimitBucket.token == 0 {
		refillSuccess = s.refill(rateLimitBucket, token)
	}

	if !refillSuccess {
		slog.Info("rate limit reached", "ipAddr", ipAddr, "pattern", pattern)
		return false
	}

	rateLimitBucket.token--

	return true
}
